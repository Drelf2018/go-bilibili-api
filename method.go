package bilibili

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Drelf2018/req"
	"github.com/Drelf2018/req/method"
)

var mixinKey string

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

// 更新 mixinKey
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
	return nil
}

// 获取 mixinKey
func GetMixinKey() (string, error) {
	if mixinKey == "" {
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
	for k, v := range query {
		if len(v) >= 1 {
			query.Set(k, unwantedChars.Replace(v[0]))
		} else {
			query.Del(k)
		}
	}
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

func (MixinKey) Query(r *http.Request, value reflect.Value, query []reflect.StructField) (err error) {
	_, err = GetMixinKey()
	if err != nil {
		return
	}
	values := method.MakeURLValues(r.Context(), value, query)
	err = AddMixinKey(values)
	if err != nil {
		return
	}
	r.URL.RawQuery = values.Encode()
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

type PostCSRF struct {
	ContentType string `req:"header" default:"application/x-www-form-urlencoded"`
}

func (PostCSRF) Method() string {
	return http.MethodPost
}

func (PostCSRF) Body(r *http.Request, value reflect.Value, body []reflect.StructField) (_ io.Reader, err error) {
	biliJct, err := r.Cookie("bili_jct")
	if err != nil {
		return nil, fmt.Errorf("%w: \"bili_jct\"", err)
	}
	if biliJct.Value == "" {
		return nil, fmt.Errorf("bilibili: empty cookie[%q]", "bili_jct")
	}
	form := method.MakeURLValues(r.Context(), value, body)
	form.Set("csrf", biliJct.Value)
	form.Set("csrf_token", biliJct.Value)
	return strings.NewReader(form.Encode()), nil
}

var _ req.APIBody = PostCSRF{}
