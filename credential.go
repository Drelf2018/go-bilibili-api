package api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/Drelf2018/req"
)

// 凭据
type Credential struct {
	// 登录 Token
	SESSDATA string

	// CSRF Token
	BiliJct string

	// 设备信息
	Buvid3 string

	// 数字型用户 UID
	DedeUserID string

	// 字符型用户 UID
	DedeUserIDckMd5 string

	// 持久化刷新口令
	// 保存在浏览器 localStorage 中 ac_time_value 的值
	RefreshToken string
}

func (c *Credential) SetCookies(_ *url.URL, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		switch cookie.Name {
		case "SESSDATA":
			c.SESSDATA = cookie.Value
		case "bili_jct":
			c.BiliJct = cookie.Value
		case "buvid3":
			c.Buvid3 = cookie.Value
		case "DedeUserID":
			c.DedeUserID = cookie.Value
		case "DedeUserID__ckMd5":
			c.DedeUserIDckMd5 = cookie.Value
		}
	}
}

func (c *Credential) Cookies(*url.URL) []*http.Cookie {
	return []*http.Cookie{
		{Name: "SESSDATA", Value: c.SESSDATA},
		{Name: "bili_jct", Value: c.BiliJct},
		{Name: "buvid3", Value: c.Buvid3},
		{Name: "DedeUserID", Value: c.DedeUserID},
		{Name: "DedeUserID__ckMd5", Value: c.DedeUserIDckMd5},
	}
}

func (c *Credential) IsValid() bool {
	return c != nil
}

func (c *Credential) UID() int {
	i, _ := strconv.Atoi(c.DedeUserID)
	return i
}

func (c *Credential) Refresh() error {
	return PostConfirmRefresh(c)
}

var _ req.CookieJar = (*Credential)(nil)
