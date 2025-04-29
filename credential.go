package api

import (
	"net/http"
	"net/url"
	"strconv"
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
	if c == nil {
		return
	}
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
	if c == nil {
		return []*http.Cookie{}
	}
	return []*http.Cookie{
		{Name: "SESSDATA", Value: c.SESSDATA},
		{Name: "bili_jct", Value: c.BiliJct},
		{Name: "buvid3", Value: c.Buvid3},
		{Name: "DedeUserID", Value: c.DedeUserID},
		{Name: "DedeUserID__ckMd5", Value: c.DedeUserIDckMd5},
	}
}

// 凭据中的用户序号
func (c *Credential) UID() int {
	i, _ := strconv.Atoi(c.DedeUserID)
	return i
}

// 判断凭据是否有效
func (c *Credential) IsValid() bool {
	result, err := GetNav(c)
	return err == nil && result.Code == 0
}

// 刷新当前凭据
func (c *Credential) Refresh() error {
	return PostConfirmRefresh(c)
}

var _ http.CookieJar = (*Credential)(nil)
