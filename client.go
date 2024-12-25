package api

import (
	"net/http"
	"net/url"

	"github.com/Drelf2018/req"
)

var cli = &req.Client{
	Header: http.Header{
		"User-Agent": {req.UserAgent},
		"Referer":    {"https://www.bilibili.com/"},
	},
}

func init() {
	var err error
	cli.BaseURL, err = url.Parse("https://api.bilibili.com/x")
	if err != nil {
		panic(err)
	}
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
