package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Drelf2018/req"
)

// 先对各个参数的名称规范化一下
// 在对查询参数的验证时可能用到两个参数名为 w_rid 和 w_webid
// 这两个参数的值在网页 js 中变量名为 mixin_key 和 access_id
// 因此本库中将以 MixinKey 和 AccessID 作为命名方式

// 检查 accessID 和 mixinKey 的间隔
//
// 每次发起请求时若用到了这两个参数，就会检查是否太久未更新，如果超出这个间隔则会自动更新一次
var CheckInterval time.Duration = time.Hour

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
	if mixinKey == "" || time.Since(mixinKeyUpdateTime) > CheckInterval {
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

// 添加请求参数验证参数
func AddMixinKey(query url.Values) error {
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
	query.Set("wts", strconv.Itoa(int(time.Now().Unix())))
	hash := md5.Sum([]byte(query.Encode() + key))
	query.Set("w_rid", hex.EncodeToString(hash[:]))
	return nil
}

// 对于需要添加 mixin_key 的接口 只需要将 MixinKey 嵌入结构体即可
type MixinKey struct{}

func (MixinKey) Query(req *http.Request, cli *req.Client, query []req.Field, value reflect.Value, api req.API) (err error) {
	// calculate mixin_key
	_, err = GetMixinKey()
	if err != nil {
		return
	}
	// create query values
	values, err := cli.MakeURLValues(query, value)
	if err != nil {
		return
	}
	// reset query
	err = AddMixinKey(values)
	if err != nil {
		return
	}
	req.URL.RawQuery = values.Encode()
	return
}

// 需要添加 mixin_key 的 GET 请求
//
// 嵌入此字段后不用额外嵌入 MixinKey 字段
type GetWBI struct{ MixinKey }

func (GetWBI) Method() string {
	return http.MethodGet
}

var _ req.APIQuery = GetWBI{}

var ErrBiliJctNotExists = errors.New("api: bili_jct does not exist")

type PostCSRF struct{}

func (PostCSRF) Method() string {
	return http.MethodPost
}

func (PostCSRF) Body(cli *req.Client, body []req.Field, value reflect.Value, api req.API) (_ io.Reader, err error) {
	// check cookie
	jar, ok := api.(http.CookieJar)
	if !ok || jar == nil {
		err = ErrBiliJctNotExists
		return
	}
	// get bili_jct
	var biliJct string
	for _, cookie := range jar.Cookies(nil) {
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
	form, err := cli.MakeURLValues(body, value)
	if err != nil {
		return
	}
	form.Set("csrf", biliJct)
	return strings.NewReader(form.Encode()), nil
}

func (PostCSRF) Header(req *http.Request, cli *req.Client, header []req.Field, value reflect.Value, api req.API) (err error) {
	err = cli.AddHeader(req, header, value)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return
}

var _ req.APIBody = PostCSRF{}
var _ req.APIHeader = PostCSRF{}
