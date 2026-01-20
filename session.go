package bilibili

import (
	"github.com/Drelf2018/req"
)

var session = req.DefaultSession.Clone().SetBaseURL("https://api.bilibili.com/x").SetHeader(map[string]string{
	"Referer": "https://www.bilibili.com/",
})

func Do[T any](api req.API) (result T, err error) {
	err = session.Result(api, &result)
	return
}

func Session() *req.Session {
	return session.Clone()
}
