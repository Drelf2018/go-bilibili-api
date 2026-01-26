package bilibili

import (
	"context"
	"net/http"
	"net/url"

	"github.com/Drelf2018/req/cookie"
)

// RefreshToken 保存在浏览器本地储存 ac_time_value 中的口令
type RefreshToken string

func (*RefreshToken) IsValid(ctx context.Context, jar http.CookieJar) (bool, error) {
	r, err := GetNav(ctx, jar)
	if err != nil || r.Code != 0 {
		return false, err
	}
	info, err := GetCookieInfo(ctx, jar)
	return !info.Data.Refresh, err
}

func (r *RefreshToken) Refresh(ctx context.Context, jar http.CookieJar) error {
	result, err := PostCookieRefresh(ctx, jar, string(*r))
	if err != nil {
		return err
	}
	err = PostConfirmRefresh(ctx, jar, string(*r))
	if err != nil {
		return err
	}
	*r = RefreshToken(result.Data.RefreshToken)
	return nil
}

var _ cookie.Refresher = (*RefreshToken)(nil)

// Credential 凭据
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

	// 持久化刷新口令，保存在浏览器 localStorage 中 ac_time_value 的值
	RefreshToken RefreshToken
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
