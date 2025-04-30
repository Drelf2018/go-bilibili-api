package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Drelf2018/req"
)

// https://socialsisteryi.github.io/bilibili-API-collect/docs/message/msg.html

// 未读消息数
type Unread struct {
	req.Get
	*Credential
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
func GetUnread(credential *Credential) (result UnreadResponse, err error) {
	err = cli.Result(Unread{Credential: credential}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/message/private_msg.html

// 未读私信数
type SingleUnread struct {
	req.Get
	*Credential

	Build      int    `api:"query,omitempty"`
	MobiApp    string `api:"query,omitempty"`
	UnreadType int    `api:"query,omitempty"`
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
func GetSingleUnread(credential *Credential) (result SingleUnreadResponse, err error) {
	err = cli.Result(SingleUnread{Credential: credential}, &result)
	return
}

// 发送私信
//
// 不知道叔叔发什么颠要求 query 和 body 都要写收发方 UID
// 猜测是为了加上 WBI 验证 因为这个验证是用 query 里的字段进行的
type SendMsg struct {
	PostCSRF
	MixinKey
	*Credential

	// 发送者 UID
	WSenderUID int `api:"query"`

	// 接收者 UID
	WReceiverID int `api:"query"`

	// 发送者 UID
	SenderUID int `api:"body" req:"msg[sender_uid]"`

	// 接收者 UID
	ReceiverID int `api:"body" req:"msg[receiver_id]"`

	// 固定为 1
	ReceiverType int `api:"body:1" req:"msg[receiver_type]"`

	// 消息类型 详见 content.go
	MsgType int `api:"body" req:"msg[msg_type]"`

	// 设备信息
	DeviceID string `api:"body:88A68CB6-CEFC-49BE-87DE-D2A22E549C1E" req:"msg[dev_id]"`

	// 秒级时间戳
	Timestamp int `api:"body" req:"msg[timestamp]"`

	// 消息内容
	Content Content `api:"body" req:"msg[content]"`

	// 未知 非必要
	MsgStatus int `api:"body:0" req:"msg[msg_status]"`

	// 表情包版本 非必要
	NewFaceVersion int `api:"body:0" req:"msg[new_face_version]"`
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
func PostSendMsg(receiver int, content Content, credential *Credential) (result SendMsgResponse, err error) {
	myUID, err := strconv.Atoi(credential.DedeUserID)
	if err != nil {
		return
	}
	err = cli.Result(SendMsg{
		WSenderUID:  myUID,
		WReceiverID: receiver,
		SenderUID:   myUID,
		ReceiverID:  receiver,
		MsgType:     content.MsgType(),
		Timestamp:   int(time.Now().Unix()),
		Content:     content,
		Credential:  credential,
	}, &result)
	return
}

// 私信消息记录
type FetchSessionMsgs struct {
	GetWBI
	*Credential

	// 聊天对象的 ID
	TalkerID int `api:"query"`

	// 聊天对象的类型 (1)用户 (2)粉丝团
	SessionType int `api:"query"`

	// 列出消息条数
	// 最大 2000
	Size int `api:"query"`

	// 发送者设备
	SenderDeviceID int `api:"query:1"`

	// 开始的序列号（开区间）
	// 默认 0 为全部
	BeginSeqno json.Number `api:"query,omitempty"`

	// 结束的序列号（开区间）
	// 默认 0 为全部
	EndSeqno json.Number `api:"query,omitempty"`
}

func (FetchSessionMsgs) RawURL() string {
	return "https://api.vc.bilibili.com/svr_sync/v1/svr_sync/fetch_session_msgs"
}

func (api *FetchSessionMsgs) ReadPage() (v FetchSessionMsgsResponse, err error) {
	err = cli.Result(api, &v)
	if err != nil {
		return
	}
	if v.Data.MinSeqno == "18446744073709551615" {
		err = ErrNoMorePage
		return
	}
	api.EndSeqno = json.Number(v.Data.MinSeqno)
	return
}

var _ PageReader[FetchSessionMsgsResponse] = (*FetchSessionMsgs)(nil)

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
func GetFetchSessionMsgs(talkerID, sessionType, size int, beginSeqno, endSeqno string, credential *Credential) (result FetchSessionMsgsResponse, err error) {
	err = cli.Result(FetchSessionMsgs{
		TalkerID:    talkerID,
		SessionType: sessionType,
		Size:        size,
		BeginSeqno:  json.Number(beginSeqno),
		EndSeqno:    json.Number(endSeqno),
		Credential:  credential,
	}, &result)
	return
}
