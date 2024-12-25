package api

import (
	"fmt"

	"github.com/Drelf2018/req"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// 注意 不应该用这个判断是否发生了错误 因为这里无法判断 Code 是否真的是错误
// 也就是说这里必然会认定为是错误 正确使用为 Error.Unwrap() 方法
func (e Error) Error() string {
	if e.Code == 0 {
		return ""
	}
	return fmt.Sprintf("api.Error: %s (%d)", e.Message, e.Code)
}

func (e Error) Unwrap() error {
	if e.Code == 0 {
		return nil
	}
	return e
}

var _ req.Unwrap = (*Error)(nil)
