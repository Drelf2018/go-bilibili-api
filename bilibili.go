package bilibili

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Drelf2018/req"
)

// Error 通用响应错误
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	caller  string
}

func (e Error) Error() string {
	if e.Code == 0 {
		return ""
	}
	if e.caller != "" {
		return fmt.Sprintf("%s: failed to request: %s (%d)", e.caller, e.Message, e.Code)
	}
	return fmt.Sprintf("bilibili: failed to request: %s (%d)", e.Message, e.Code)
}

func (e Error) Unwrap() error {
	if e.Code == 0 {
		return nil
	}
	return e
}

var _ req.Unwrap = Error{}

var Session = req.DefaultSession.Clone().SetHeader(map[string]string{
	"Referer": "https://www.bilibili.com/",
}).SetBaseURL("https://api.bilibili.com/x")

func Do[T any](ctx context.Context, api req.API) (result T, err error) {
	err = Session.ResultWithContext(ctx, api, &result)
	if err == nil {
		return
	}
	e, ok := err.(Error)
	if !ok {
		return
	}
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return
	}
	funcObj := runtime.FuncForPC(pc)
	if funcObj == nil {
		return
	}
	e.caller = filepath.Base(funcObj.Name())
	return result, e
}
