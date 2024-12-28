package api

import "github.com/Drelf2018/req"

// https://socialsisteryi.github.io/bilibili-API-collect/docs/live/info.html

// 获取直播间信息
type RoomInfo struct {
	req.Get

	// 直播间号	可以为短号
	RoomID int `api:"query"`
}

func (RoomInfo) RawURL() string {
	return "https://api.live.bilibili.com/room/v1/Room/get_info"
}

type RoomInfoResponse struct {
	Error
	Msg  string `json:"msg"` // "ok"
	Data struct {
		UID              int      `json:"uid"`         // 434334701
		RoomID           int      `json:"room_id"`     // 21452505
		ShortID          int      `json:"short_id"`    // 0
		Attention        int      `json:"attention"`   // 1057100
		Online           int      `json:"online"`      // 0
		IsPortrait       bool     `json:"is_portrait"` // false
		Description      string   `json:"description"`
		LiveStatus       int      `json:"live_status"`        // 0
		AreaID           int      `json:"area_id"`            // 745
		ParentAreaID     int      `json:"parent_area_id"`     // 9
		ParentAreaName   string   `json:"parent_area_name"`   // "虚拟主播"
		OldAreaID        int      `json:"old_area_id"`        // 6
		Background       string   `json:"background"`         // "https://i0.hdslb.com/bfs/live/636d66a97d5f55099a9d8d6813558d6d4c95fd61.jpg"
		Title            string   `json:"title"`              // "看看大家的24年最值购物"
		UserCover        string   `json:"user_cover"`         // "https://i0.hdslb.com/bfs/live/new_room_cover/ef9375fd23aefb03e3c2fd48934fa51a30caacbc.jpg"
		Keyframe         string   `json:"keyframe"`           // ""
		IsStrictRoom     bool     `json:"is_strict_room"`     // false
		LiveTime         string   `json:"live_time"`          // "0000-00-00 00:00:00"
		Tags             string   `json:"tags"`               // "七海,海子姐,VirtuaReal"
		IsAnchor         int      `json:"is_anchor"`          // 0
		RoomSilentType   string   `json:"room_silent_type"`   // ""
		RoomSilentLevel  int      `json:"room_silent_level"`  // 0
		RoomSilentSecond int      `json:"room_silent_second"` // 0
		AreaName         string   `json:"area_name"`          // "虚拟Gamer"
		Pendants         string   `json:"pendants"`           // ""
		AreaPendants     string   `json:"area_pendants"`      // ""
		HotWords         []string `json:"hot_words"`
		HotWordsStatus   int      `json:"hot_words_status"` // 0
		Verify           string   `json:"verify"`           // ""
		NewPendants      struct {
			Frame struct {
				Name       string `json:"name"`         // "长红计划-Topstar"
				Value      string `json:"value"`        // "https://i0.hdslb.com/bfs/live/62e62f657d379aaaec2bbd4a6a16a938bcba76e6.png"
				Position   int    `json:"position"`     // 0
				Desc       string `json:"desc"`         // ""
				Area       int    `json:"area"`         // 0
				AreaOld    int    `json:"area_old"`     // 0
				BgColor    string `json:"bg_color"`     // ""
				BgPic      string `json:"bg_pic"`       // ""
				UseOldArea bool   `json:"use_old_area"` // false
			} `json:"frame"`
			Badge struct {
				Name     string `json:"name"`     // "v_person"
				Position int    `json:"position"` // 3
				Value    string `json:"value"`    // ""
				Desc     string `json:"desc"`     // "虚拟UP主、bilibili直播高能主播"
			} `json:"badge"`
			MobileFrame struct {
				Name       string `json:"name"`         // "长红计划-Topstar"
				Value      string `json:"value"`        // "https://i0.hdslb.com/bfs/live/62e62f657d379aaaec2bbd4a6a16a938bcba76e6.png"
				Position   int    `json:"position"`     // 0
				Desc       string `json:"desc"`         // ""
				Area       int    `json:"area"`         // 0
				AreaOld    int    `json:"area_old"`     // 0
				BgColor    string `json:"bg_color"`     // ""
				BgPic      string `json:"bg_pic"`       // ""
				UseOldArea bool   `json:"use_old_area"` // false
			} `json:"mobile_frame"`
			MobileBadge any `json:"mobile_badge"`
		} `json:"new_pendants"`
		UpSession            string `json:"up_session"`              // ""
		PkStatus             int    `json:"pk_status"`               // 0
		PkID                 int    `json:"pk_id"`                   // 0
		BattleID             int    `json:"battle_id"`               // 0
		AllowChangeAreaTime  int    `json:"allow_change_area_time"`  // 0
		AllowUploadCoverTime int    `json:"allow_upload_cover_time"` // 0
		StudioInfo           struct {
			Status     int   `json:"status"` // 0
			MasterList []any `json:"master_list"`
		} `json:"studio_info"`
	} `json:"data"`
}

// 获取直播间信息
func GetRoomInfo(roomid int) (result RoomInfoResponse, err error) {
	err = cli.Result(RoomInfo{RoomID: roomid}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/live/message_stream.html

// 获取信息流认证密钥
type DanmuInfo struct {
	req.Get
	*Credential

	// 直播间真实 ID
	ID int `api:"query"`
}

func (DanmuInfo) RawURL() string {
	return "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo"
}

type DanmuInfoResponse struct {
	Error
	Data struct {
		Group            string  `json:"group"`              // "live"
		BusinessID       int     `json:"business_id"`        // 0
		RefreshRowFactor float64 `json:"refresh_row_factor"` // 0.125
		RefreshRate      int     `json:"refresh_rate"`       // 100
		MaxDelay         int     `json:"max_delay"`          // 5000
		Token            string  `json:"token"`              // "kCFEQiWeDA15-4KOmwu4IZ597OPNSEfHoYTGY0HOoWQEHVBKM4nGTGOWW6SLE2LJxnOKZINMCQxIExb7xkO8FIBZNs_kMflWYeojYTEV_jl9TsILOJRX1pC2kAOByp72j4WrKsc0GdvdUnxqwfnib8Ysps1Upwe2O2Db_uRSr6hVKEAAlwjifVvc"
		HostList         []struct {
			Host    string `json:"host"`     // "zj-cn-live-comet.chat.bilibili.com"
			Port    int    `json:"port"`     // 2243
			WssPort int    `json:"wss_port"` // 2245
			WsPort  int    `json:"ws_port"`  // 2244
		} `json:"host_list"`
	} `json:"data"`
}

// 获取信息流认证密钥
//
// 参数 roomid 为直播间真实 ID
func GetDanmuInfo(roomid int, credential *Credential) (result DanmuInfoResponse, err error) {
	err = cli.Result(DanmuInfo{ID: roomid, Credential: credential}, &result)
	return
}
