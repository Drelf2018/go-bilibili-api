package api

import (
	"net/http"
	"net/url"

	"github.com/Drelf2018/req"

	_ "unsafe"
)

//go:linkname cli
var cli = &req.Client{
	Header: http.Header{
		"User-Agent": {req.UserAgent},
		"Referer":    {"https://www.bilibili.com/"},
	},
}

func init() {
	cli.BaseURL, _ = url.Parse("https://api.bilibili.com/x")
}

func Do[T any](api req.API) (result T, err error) {
	err = cli.Result(api, &result)
	return
}

func GetClient() *req.Client {
	return &req.Client{
		BaseURL: cli.BaseURL,
		Header:  cli.Header.Clone(),
	}
}
