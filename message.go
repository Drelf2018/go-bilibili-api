package bilibili

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Drelf2018/req"
)

// https://socialsisteryi.github.io/bilibili-API-collect/docs/message/msg.html

// 未读消息数
type Unread struct {
	req.Get
	http.CookieJar
}

func (Unread) RawURL() string {
	return "https://api.vc.bilibili.com/x/im/web/msgfeed/unread"
}

type UnreadResponse struct {
	Error
	Data struct {
		At          int `json:"at"`
		Chat        int `json:"chat"`
		Coin        int `json:"coin"`
		Danmu       int `json:"danmu"`
		Favorite    int `json:"favorite"`
		Like        int `json:"like"`
		RecvLike    int `json:"recv_like"`
		RecvReply   int `json:"recv_reply"`
		Reply       int `json:"reply"`
		SysMsg      int `json:"sys_msg"`
		SysMsgStyle int `json:"sys_msg_style"`
		Up          int `json:"up"`
	} `json:"data"`
}

// 未读消息数
func GetUnread(ctx context.Context, jar http.CookieJar) (UnreadResponse, error) {
	return Do[UnreadResponse](ctx, Unread{CookieJar: jar})
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/message/private_msg.html

// 未读私信数
type SingleUnread struct {
	req.Get
	http.CookieJar

	Build      int    `req:"query,omitempty"`
	MobiApp    string `req:"query,omitempty"`
	UnreadType int    `req:"query,omitempty"`
}

func (SingleUnread) RawURL() string {
	return "https://api.vc.bilibili.com/session_svr/v1/session_svr/single_unread"
}

type SingleUnreadResponse struct {
	Error
	Msg  string `json:"msg"`
	Data struct {
		UnfollowUnread       int `json:"unfollow_unread"`
		FollowUnread         int `json:"follow_unread"`
		UnfollowPushMsg      int `json:"unfollow_push_msg"`
		DustbinPushMsg       int `json:"dustbin_push_msg"`
		DustbinUnread        int `json:"dustbin_unread"`
		BizMsgUnfollowUnread int `json:"biz_msg_unfollow_unread"`
		BizMsgFollowUnread   int `json:"biz_msg_follow_unread"`
		CustomUnread         int `json:"custom_unread"`
	} `json:"data"`
}

// 未读私信数
func GetSingleUnread(ctx context.Context, jar http.CookieJar) (SingleUnreadResponse, error) {
	return Do[SingleUnreadResponse](ctx, SingleUnread{CookieJar: jar})
}

// 发送私信
//
// 不知道叔叔发什么颠要求 query 和 body 都要写收发方 UID
// 猜测是为了加上 WBI 验证 因为这个验证是用 query 里的字段进行的
type SendMsg struct {
	PostCSRF
	MixinKey
	http.CookieJar

	// 发送者 UID
	WSenderUID int `req:"query"`

	// 接收者 UID
	WReceiverID int `req:"query"`

	// 发送者 UID
	SenderUID int `req:"body:msg[sender_uid]"`

	// 接收者 UID
	ReceiverID int `req:"body:msg[receiver_id]"`

	// 固定为 1
	ReceiverType int `req:"body" default:"1:msg[receiver_type]"`

	// 消息类型 详见 content.go
	MsgType MsgType `req:"body:msg[msg_type]"`

	// 设备信息
	DeviceID string `req:"body" default:"88A68CB6-CEFC-49BE-87DE-D2A22E549C1E:msg[dev_id]"`

	// 秒级时间戳
	Timestamp int `req:"body:msg[timestamp]"`

	// 消息内容
	Content Content `req:"body:msg[content]"`

	// 未知 非必要
	MsgStatus int `req:"body" default:"0:msg[msg_status]"`

	// 表情包版本 非必要
	NewFaceVersion int `req:"body" default:"0:msg[new_face_version]"`
}

func (SendMsg) RawURL() string {
	return "https://api.vc.bilibili.com/web_im/v1/web_im/send_msg"
}

type SendMsgResponse struct {
	Error
	Data struct {
		MsgKey int64 `json:"msg_key"`
		EInfos []struct {
			Text string `json:"text"`
			URL  string `json:"url"`
			Size int    `json:"size"`
		} `json:"e_infos"`
		MsgContent  string `json:"msg_content"`
		KeyHitInfos struct {
		} `json:"key_hit_infos"`
	} `json:"data"`
}

// 发送私信
func PostSendMsg(ctx context.Context, sender int, receiver int, content Content, jar http.CookieJar) (SendMsgResponse, error) {
	return Do[SendMsgResponse](ctx, SendMsg{
		WSenderUID:  sender,
		WReceiverID: receiver,
		SenderUID:   sender,
		ReceiverID:  receiver,
		MsgType:     content.MsgType(),
		Timestamp:   int(time.Now().Unix()),
		Content:     content,
		CookieJar:   jar,
	})
}

// 私信消息记录
type FetchSessionMsgs struct {
	GetWBI
	http.CookieJar

	// 聊天对象的 ID
	TalkerID int `req:"query"`

	// 聊天对象的类型 (1)用户 (2)粉丝团
	SessionType int `req:"query"`

	// 列出消息条数
	// 最大 2000
	Size int `req:"query"`

	// 发送者设备
	SenderDeviceID int `req:"query" default:"1"`

	// 开始的序列号（开区间）
	// 默认 0 为全部
	BeginSeqno json.Number `req:"query,omitempty"`

	// 结束的序列号（开区间）
	// 默认 0 为全部
	EndSeqno json.Number `req:"query,omitempty"`
}

func (FetchSessionMsgs) RawURL() string {
	return "https://api.vc.bilibili.com/svr_sync/v1/svr_sync/fetch_session_msgs"
}

type Message struct {
	SenderUID      int         `json:"sender_uid"`
	ReceiverType   int         `json:"receiver_type"`
	ReceiverID     int         `json:"receiver_id"`
	MsgType        int         `json:"msg_type"`
	Content        string      `json:"content"`
	MsgSeqno       json.Number `json:"msg_seqno"`
	Timestamp      int         `json:"timestamp"`
	AtUIDs         []int       `json:"at_uids" gorm:"serializer:json"`
	MsgKey         int64       `json:"msg_key"`
	MsgStatus      int         `json:"msg_status"`
	NotifyCode     string      `json:"notify_code"`
	NewFaceVersion int         `json:"new_face_version,omitempty"`
}

func (m Message) String() string {
	return fmt.Sprintf("%d: %s (%s)", m.SenderUID, m.Content, m.MsgSeqno)
}

type FetchSessionMsgsResponse struct {
	Error
	Msg  string `json:"msg"`
	Data struct {
		Messages []Message   `json:"messages"`
		HasMore  int         `json:"has_more"`
		MinSeqno json.Number `json:"min_seqno"`
		MaxSeqno json.Number `json:"max_seqno"`
		EInfos   []struct {
			Text string `json:"text"`
			URL  string `json:"url"`
			Size int    `json:"size"`
		} `json:"e_infos"`
	} `json:"data"`
}

// 私信消息记录
//
// talkerID: 聊天对象的 ID
//
// sessionType: 聊天对象的类型 (1)用户 (2)粉丝团
//
// size: 列出消息条数 最大 2000
//
// beginSeqno: 消息开始的序列号（开区间） 默认 0 为全部
//
// endSeqno: 消息结束的序列号（开区间） 默认 0 为全部
func GetFetchSessionMsgs(ctx context.Context, talkerID, sessionType, size int, beginSeqno, endSeqno string, jar http.CookieJar) (FetchSessionMsgsResponse, error) {
	return Do[FetchSessionMsgsResponse](ctx, FetchSessionMsgs{
		TalkerID:    talkerID,
		SessionType: sessionType,
		Size:        size,
		BeginSeqno:  json.Number(beginSeqno),
		EndSeqno:    json.Number(endSeqno),
		CookieJar:   jar,
	})
}
