package bilibili

import (
	"encoding/json"

	"github.com/Drelf2018/req/method"
)

// MsgType 表示消息类型的整数
type MsgType int

const (
	MsgText        = 1   // 纯文字消息
	MsgImage       = 2   // 图片消息
	MsgWithdraw    = 5   // 撤回消息
	MsgGroupsImage = 6   // 应援团图片
	MsgShareVideo  = 7   // 分享视频
	MsgNotice      = 10  // 系统通知
	MsgPushedVideo = 11  // UP 主推送的视频
	MsgWelcome     = 306 // 新成员加入应援团欢迎
)

// Content 内容接口
type Content interface {
	MsgType() MsgType
}

// Text 纯文本内容
type Text string

func (Text) MsgType() MsgType { return MsgText }

var _ Content = Text("")

func (t Text) MarshalString() string {
	b, _ := json.Marshal(map[string]Text{"content": t})
	return string(b)
}

var _ method.Marshaler = Text("")

// Image 图片内容
type Image struct {
	URL      string `json:"url"`                // 图片链接 默认为相簿图片上传通道 也可用三方图床
	Width    int    `json:"width,omitempty"`    // 图片的宽	单位像素 非必要
	Height   int    `json:"height,omitempty"`   // 图片的高	单位像素 非必要
	Type     string `json:"type,omitempty"`     // 图片格式	非必要
	Original int    `json:"original,omitempty"` // 作用未知 默认值 1 非必要
	Size     int    `json:"size,omitempty"`     // 文件大小	单位千字节 非必要
}

func (Image) MsgType() MsgType { return MsgImage }

var _ Content = Image{}

// Withdraw 撤回消息内容
type Withdraw string

func (Withdraw) MsgType() MsgType { return MsgWithdraw }

var _ Content = Withdraw("")
