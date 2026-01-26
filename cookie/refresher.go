package cookie

import (
	"net/http"
	"time"

	bilibili "github.com/Drelf2018/go-bilibili-api"
	"gorm.io/gorm"
)

// Cookie 是用于 GORM 存取的 Cookie 模型
type Cookie struct {
	Name        string
	Value       string
	Quoted      bool
	CreatedAt   time.Time
	RefresherID uint64
}

// Refresher 用保存在浏览器本地储存 ac_time_value 中的口令实现持久化刷新
type Refresher struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement"`
	Jar       http.CookieJar `gorm:"-"`
	Token     bilibili.RefreshToken
	Cookies   []Cookie
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// ToHTTPCookies 将 GORM 模型切片转换成 *http.Cookie 切片
func ToHTTPCookies(models []Cookie) []*http.Cookie {
	cookies := make([]*http.Cookie, 0, len(models))
	for _, m := range models {
		cookies = append(cookies, &http.Cookie{Name: m.Name, Value: m.Value, Quoted: m.Quoted})
	}
	return cookies
}

// WriteTo 将刷新器中的 Cookies 写入 http.CookieJar
func (r *Refresher) WriteTo(jar http.CookieJar) {
	if jar != nil {
		jar.SetCookies(bilibili.Session.BaseURL, ToHTTPCookies(r.Cookies))
	}
}

// AfterFind 在查询后将数据写入 http.CookieJar
func (r *Refresher) AfterFind(*gorm.DB) error {
	if r.Jar == nil {
		r.Jar = &bilibili.Credential{RefreshToken: r.Token}
	}
	r.WriteTo(r.Jar)
	return nil
}

// FromHTTPCookies 将 *http.Cookie 切片转换成 GORM 模型切片
func FromHTTPCookies(cookies []*http.Cookie) []Cookie {
	models := make([]Cookie, 0, len(cookies))
	for _, cookie := range cookies {
		if cookie.Name != "" {
			models = append(models, Cookie{Name: cookie.Name, Value: cookie.Value, Quoted: cookie.Quoted})
		}
	}
	return models
}

// ReadFrom 从 http.CookieJar 中读取 Cookies 到刷新器
func (r *Refresher) ReadFrom(jar http.CookieJar) {
	if jar != nil {
		r.Cookies = append(r.Cookies, FromHTTPCookies(jar.Cookies(bilibili.Session.BaseURL))...)
	}
}

// BeforeSave 在保存前读取 http.CookieJar 中数据
func (r *Refresher) BeforeSave(*gorm.DB) error {
	if r.Jar != nil {
		r.ReadFrom(r.Jar)
	}
	return nil
}
