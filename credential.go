package bilibili

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Drelf2018/req/cookie"
)

// 凭据
type Credential struct {
	// 登录 Token
	SESSDATA string `cookie:"SESSDATA"`

	// CSRF Token
	BiliJct string `cookie:"bili_jct"`

	// 设备信息
	Buvid3 string `cookie:"buvid3"`

	// 数字型用户 UID
	DedeUserID string `cookie:"DedeUserID"`

	// 字符型用户 UID
	DedeUserIDckMd5 string `cookie:"DedeUserID__ckMd5"`

	// 持久化刷新口令
	// 保存在浏览器 localStorage 中 ac_time_value 的值
	RefreshToken string
}

func (c *Credential) SetCookies(_ *url.URL, cookies []*http.Cookie) {
	if c != nil {
		cookie.Set(c, cookies)
	}
}

func (c *Credential) Cookies(*url.URL) (cookies []*http.Cookie) {
	if c != nil {
		cookies, _ = cookie.Get(c)
	}
	return
}

var _ http.CookieJar = (*Credential)(nil)

// 凭据中的用户标识符
func (c *Credential) UID() int {
	i, _ := strconv.Atoi(c.DedeUserID)
	return i
}

// Cookie 是用于 GORM 存取的 Cookie 模型
type Cookie struct {
	Name        string
	Value       string
	CreatedAt   time.Time
	RefresherID uint64
}

// Refresher 用保存在浏览器本地储存 ac_time_value 中的口令实现持久化刷新
type Refresher struct {
	ID        uint64 `gorm:"primarykey"`
	Token     string
	Cookies   []Cookie
	CreatedAt time.Time
}

func (*Refresher) IsValid(ctx context.Context, jar http.CookieJar) (bool, error) {
	r, err := GetNavWithContext(ctx, jar)
	if err != nil || r.Code != 0 {
		return false, err
	}
	info, err := GetCookieInfoWithContext(ctx, jar)
	return !info.Data.Refresh, err
}

func (r *Refresher) Refresh(ctx context.Context, jar http.CookieJar) (err error) {
	r.Token, err = PostConfirmRefreshWithContext(ctx, r.Token, jar)
	return
}

var _ cookie.Refresher = (*Refresher)(nil)

// FromHTTPCookies 将 *http.Cookie 切片转换成 GORM 模型切片
func FromHTTPCookies(cookies []*http.Cookie) []Cookie {
	models := make([]Cookie, 0, len(cookies))
	for _, cookie := range cookies {
		if cookie.Name != "" {
			models = append(models, Cookie{Name: cookie.Name, Value: cookie.Value})
		}
	}
	return models
}

// ToHTTPCookies 将 GORM 模型切片转换成 *http.Cookie 切片
func ToHTTPCookies(models []Cookie) []*http.Cookie {
	cookies := make([]*http.Cookie, 0, len(models))
	for _, m := range models {
		cookies = append(cookies, &http.Cookie{Name: m.Name, Value: m.Value})
	}
	return cookies
}
