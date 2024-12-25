package api

import (
	"encoding/json"

	"github.com/Drelf2018/req"
)

const (
	// 纯文字消息
	MsgTypeText = 1

	// 图片消息
	MsgTypeImage = 2

	// 撤回消息
	MsgTypeWithdraw = 5

	// 应援团图片
	// 但似乎不常触发 一般使用 MsgTypeImage 即可
	MsgTypeGroupsImage = 6

	// 分享视频
	MsgTypeShareVideo = 7

	// 系统通知
	MsgTypeNotice = 10

	// UP 主推送的视频
	MsgTypePushedVideo = 11

	// 新成员加入应援团欢迎
	MsgTypeWelcome = 306
)

// 内容接口
type Content interface {
	req.Marshaler
	MsgType() int
}

// 纯文本内容
type TextContent string

func (TextContent) MsgType() int { return MsgTypeText }

func (c TextContent) MarshalString() (string, error) {
	b, err := json.Marshal(map[string]TextContent{"content": c})
	if err != nil {
		return "", err
	}
	return string(b), nil
}

var _ Content = TextContent("")

// 图片内容
type ImageContent struct {
	URL      string `json:"url"`                // 图片链接 默认为相簿图片上传通道 也可用三方图床
	Width    int    `json:"width,omitempty"`    // 图片的宽	单位像素 非必要
	Height   int    `json:"height,omitempty"`   // 图片的高	单位像素 非必要
	Type     string `json:"type,omitempty"`     // 图片格式	非必要
	Original int    `json:"original,omitempty"` // 作用未知 默认值 1 非必要
	Size     int    `json:"size,omitempty"`     // 文件大小	单位千字节 非必要
}

func (ImageContent) MsgType() int { return MsgTypeImage }

func (c ImageContent) MarshalString() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

var _ Content = (*ImageContent)(nil)

// 撤回消息内容
type WithdrawContent string

func (WithdrawContent) MsgType() int { return MsgTypeWithdraw }

func (c WithdrawContent) MarshalString() (string, error) {
	return string(c), nil
}

var _ Content = WithdrawContent("")
