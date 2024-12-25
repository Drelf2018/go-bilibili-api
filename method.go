package api

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Drelf2018/req"
)

// 检查 accessID 和 mixinKey 的间隔
//
// 每次发起请求时若用到了这两个参数，就会检查是否太久未更新，如果超出这个间隔则会自动更新一次
var CheckInterval time.Duration = 10 * time.Minute

var accessID string

var accessIDUpdateTime time.Time

var accessPattern = regexp.MustCompile(`<script id="__RENDER_DATA__" type="application/json">(.*?)</script>`)

var ErrAccessNotExist = errors.New("api: access_id does not exist")

// 更新 accessID（一般不需要手动调用）
//
// 成功后会刷新 accessIDUpdateTime
func UpdateAccessID() error {
	resp, err := http.Get("https://space.bilibili.com/208259/dynamic")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r := accessPattern.FindStringSubmatch(string(b))
	if len(r) < 2 {
		return ErrAccessNotExist
	}

	s, err := url.QueryUnescape(r[1])
	if err != nil {
		return err
	}

	var data struct {
		AccessID string `json:"access_id"`
	}
	err = json.Unmarshal([]byte(s), &data)
	if err != nil {
		return err
	}

	accessID = data.AccessID
	accessIDUpdateTime = time.Now()
	return nil
}

// 获取 accessID
//
// 无法保证一定可用，因为可能在上次检查后服务端就刷新了
func GetAccessID() (string, error) {
	if accessID == "" || time.Since(accessIDUpdateTime) > CheckInterval || time.Now().Day() != accessIDUpdateTime.Day() {
		err := UpdateAccessID()
		if err != nil {
			return "", err
		}
	}
	return accessID, nil
}

type NeedAccess interface {
	NeedAccess()
}

// 对于需要添加 access_id 的接口 只需要将 AccessID 嵌入结构体即可
type AccessID struct{}

func (AccessID) NeedAccess() {}

var mixinKey string

var mixinKeyUpdateTime time.Time

// 申必魔法数字列表
var mixinKeyEncTab = []int{
	46, 47, 18, 2, 53, 8, 23, 32,
	15, 50, 10, 31, 58, 3, 45, 35,
	27, 43, 5, 49, 33, 9, 42, 19,
	29, 28, 14, 39, 12, 38, 41, 13,
	37, 48, 7, 16, 24, 55, 40, 61,
	26, 17, 0, 1, 60, 51, 30, 4,
	22, 25, 54, 21, 56, 59, 6, 63,
	57, 62, 11, 36, 20, 34, 44, 52,
}

// 根据拼接后字符串生成 mixinKey
func GenerateMixinKey(ae string) string {
	buf := &bytes.Buffer{}
	for _, val := range mixinKeyEncTab {
		if val < len(ae) {
			buf.WriteByte(ae[val])
		}
	}
	return buf.String()[:32]
}

// 提取 img_key 和 sub_key
func SplitURL(url string) string {
	return strings.Replace(filepath.Base(url), filepath.Ext(url), "", 1)
}

// 更新 mixinKey（一般不需要手动调用）
//
// 成功后会刷新 mixinKeyUpdateTime
func UpdateMixinKey() error {
	r, err := GetNav(nil)
	if err != nil {
		return err
	}
	if r.Code != 0 && r.Code != -101 {
		e := Error{r.Code, r.Message}
		return e.Unwrap()
	}
	mixinKey = GenerateMixinKey(SplitURL(r.Data.WbiImg.ImgURL) + SplitURL(r.Data.WbiImg.SubURL))
	mixinKeyUpdateTime = time.Now()
	return nil
}

// 获取 mixinKey
//
// 无法保证一定可用，因为可能在上次检查后服务端就刷新了
func GetMixinKey() (string, error) {
	if mixinKey == "" || time.Since(mixinKeyUpdateTime) > CheckInterval || time.Now().Day() != mixinKeyUpdateTime.Day() {
		err := UpdateMixinKey()
		if err != nil {
			return "", err
		}
	}
	return mixinKey, nil
}

var unwantedChars = strings.NewReplacer(
	"!", "",
	"'", "",
	"(", "",
	")", "",
	"*", "",
)

// 加密请求参数
func EncodeWBI(query url.Values) error {
	// remove unwanted characters
	for k, v := range query {
		if len(v) >= 1 {
			query.Set(k, unwantedChars.Replace(v[0]))
		} else {
			query.Del(k)
		}
	}
	// reset query
	key, err := GetMixinKey()
	if err != nil {
		return err
	}
	query.Del("w_rid")
	query.Set("web_location", "333.999")
	query.Set("wts", strconv.Itoa(int(time.Now().Unix())))
	hash := md5.Sum([]byte(query.Encode() + key))
	query.Set("w_rid", hex.EncodeToString(hash[:]))
	return nil
}

type NeedWBI interface {
	NeedWBI()
}

// 对于需要添加 mixin_key 的接口 只需要将 WBI 嵌入结构体即可
type WBI struct{}

func (WBI) NeedWBI() {}

// 需要添加 mixin_key 的 GET 请求
//
// 嵌入此字段后不用再嵌入 WBI 字段
type GetWBI struct{}

func (GetWBI) Method() string {
	return http.MethodGet
}

func (GetWBI) NewRequestWithContext(ctx context.Context, cli *req.Client, api req.APIData) (r *http.Request, err error) {
	// calculate mixin_key
	_, err = GetMixinKey()
	if err != nil {
		return
	}
	// calculate access_id
	if _, ok := api.(NeedAccess); ok {
		_, err = GetAccessID()
		if err != nil {
			return
		}
	}
	// create request
	r, err = cli.AddBody(ctx, api, nil)
	if err != nil {
		return
	}
	// create query values
	task := req.LoadTask(api)
	value := reflect.Indirect(reflect.ValueOf(api))
	query, err := cli.MakeURLValues(task.Query, value)
	if err != nil {
		return
	}
	// reset query
	if _, ok := api.(NeedAccess); ok {
		query.Set("w_webid", accessID)
	}
	err = EncodeWBI(query)
	if err != nil {
		return
	}
	r.URL.RawQuery = query.Encode()
	// add header
	err = cli.AddHeader(r, task.Header, value)
	return
}

var _ req.APICreator = GetWBI{}

var ErrBiliJctNotExists = errors.New("api: bili_jct does not exist")

type PostCSRF struct{}

func (PostCSRF) Method() string {
	return http.MethodPost
}

func (PostCSRF) NewRequestWithContext(ctx context.Context, cli *req.Client, api req.APIData) (r *http.Request, err error) {
	// check cookie
	jar, ok := api.(req.CookieJar)
	if !ok || !jar.IsValid() {
		err = ErrBiliJctNotExists
		return
	}
	// generate request url
	var u *url.URL
	rawURL := api.RawURL()
	if cli.BaseURL != nil && strings.HasPrefix(rawURL, "/") {
		u = cli.BaseURL.JoinPath(rawURL)
	} else {
		u, err = url.Parse(rawURL)
		if err != nil {
			return
		}
	}
	// get bili_jct
	var biliJct string
	for _, cookie := range jar.Cookies(u) {
		if cookie.Name == "bili_jct" {
			biliJct = cookie.Value
			break
		}
	}
	if biliJct == "" {
		err = ErrBiliJctNotExists
		return
	}
	// create form body
	task := req.LoadTask(api)
	value := reflect.Indirect(reflect.ValueOf(api))
	form := url.Values{"csrf": {biliJct}} //, "csrf_token": {biliJct}
	for _, data := range task.Body {
		err = cli.AddValue(form, data, value)
		if err != nil {
			return
		}
	}
	// create request
	r, err = cli.AddBody(ctx, api, strings.NewReader(form.Encode()))
	if err != nil {
		return
	}
	// wbi verify
	if _, ok := api.(NeedWBI); ok {
		var query url.Values
		query, err = cli.MakeURLValues(task.Query, value)
		if err != nil {
			return
		}
		err = EncodeWBI(query)
		if err != nil {
			return
		}
		r.URL.RawQuery = query.Encode()
	} else {
		err = cli.AddQuery(r, task.Query, value)
		if err != nil {
			return
		}
	}
	// add header
	err = cli.AddHeader(r, task.Header, value)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return
}

var _ req.APICreator = PostCSRF{}
