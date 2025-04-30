package api

import (
	"strconv"
	"strings"

	"github.com/Drelf2018/req"
)

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/info.html

// 用户空间详细信息
type AccInfo struct {
	GetWBI
	*Credential

	// 目标用户 UID
	MID int `api:"query"`
}

func (AccInfo) RawURL() string {
	return "/space/wbi/acc/info"
}

type AccInfoResponse struct {
	Error
	Data struct {
		MID         int    `json:"mid"`           // 208259
		Name        string `json:"name"`          // "陈睿"
		Sex         string `json:"sex"`           // "男"
		Face        string `json:"face"`          // "http://i2.hdslb.com/bfs/app/8920e6741fc2808cce5b81bc27abdbda291655d3.png"
		FaceNft     int    `json:"face_nft"`      // 0
		FaceNftType int    `json:"face_nft_type"` // 0
		Sign        string `json:"sign"`          // "喜欢的话就坚持吧"
		Rank        int    `json:"rank"`          // 10000
		Level       int    `json:"level"`         // 6
		Jointime    int    `json:"jointime"`      // 0
		Moral       int    `json:"moral"`         // 0
		Silence     int    `json:"silence"`       // 0
		Coins       int    `json:"coins"`         // 0
		FansBadge   bool   `json:"fans_badge"`    // true
		FansMedal   struct {
			Show  bool `json:"show"` // false
			Wear  bool `json:"wear"` // false
			Medal any  `json:"medal"`
		} `json:"fans_medal"`
		Official struct {
			Role  int    `json:"role"`  // 2
			Title string `json:"title"` // "bilibili董事长兼CEO"
			Desc  string `json:"desc"`  // ""
			Type  int    `json:"type"`  // 0
		} `json:"official"`
		Vip struct {
			Type       int   `json:"type"`         // 2
			Status     int   `json:"status"`       // 1
			DueDate    int64 `json:"due_date"`     // 2096121600000
			VipPayType int   `json:"vip_pay_type"` // 1
			ThemeType  int   `json:"theme_type"`   // 0
			Label      struct {
				Path                  string `json:"path"`                      // ""
				Text                  string `json:"text"`                      // "十年大会员"
				LabelTheme            string `json:"label_theme"`               // "ten_annual_vip"
				TextColor             string `json:"text_color"`                // "#FFFFFF"
				BgStyle               int    `json:"bg_style"`                  // 1
				BgColor               string `json:"bg_color"`                  // "#FB7299"
				BorderColor           string `json:"border_color"`              // ""
				UseImgLabel           bool   `json:"use_img_label"`             // true
				ImgLabelURIHans       string `json:"img_label_uri_hans"`        // "https://i0.hdslb.com/bfs/activity-plat/static/20220608/e369244d0b14644f5e1a06431e22a4d5/wltavwHAkL.gif"
				ImgLabelURIHant       string `json:"img_label_uri_hant"`        // ""
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"` // "https://i0.hdslb.com/bfs/vip/802418ff03911645648b63aa193ba67997b5a0bc.png"
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"` // "https://i0.hdslb.com/bfs/activity-plat/static/20220614/e369244d0b14644f5e1a06431e22a4d5/8u7iRTPE7N.png"
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`     // 1
			NicknameColor      string `json:"nickname_color"`       // "#FB7299"
			Role               int    `json:"role"`                 // 7
			AvatarSubscriptURL string `json:"avatar_subscript_url"` // ""
			TvVipStatus        int    `json:"tv_vip_status"`        // 1
			TvVipPayType       int    `json:"tv_vip_pay_type"`      // 0
			TvDueDate          int    `json:"tv_due_date"`          // 1873641600
			AvatarIcon         struct {
				IconType     int `json:"icon_type"` // 1
				IconResource struct {
				} `json:"icon_resource"`
			} `json:"avatar_icon"`
		} `json:"vip"`
		Pendant struct {
			PID               int    `json:"pid"`                 // 0
			Name              string `json:"name"`                // ""
			Image             string `json:"image"`               // ""
			Expire            int    `json:"expire"`              // 0
			ImageEnhance      string `json:"image_enhance"`       // ""
			ImageEnhanceFrame string `json:"image_enhance_frame"` // ""
			NPID              int    `json:"n_pid"`               // 0
		} `json:"pendant"`
		Nameplate struct {
			NID        int    `json:"nid"`         // 0
			Name       string `json:"name"`        // ""
			Image      string `json:"image"`       // ""
			ImageSmall string `json:"image_small"` // ""
			Level      string `json:"level"`       // ""
			Condition  string `json:"condition"`   // ""
		} `json:"nameplate"`
		UserHonourInfo struct {
			MID               int   `json:"mid"` // 0
			Colour            any   `json:"colour"`
			Tags              []any `json:"tags"`
			IsLatest100Honour int   `json:"is_latest_100honour"` // 0
		} `json:"user_honour_info"`
		IsFollowed bool   `json:"is_followed"` // false
		TopPhoto   string `json:"top_photo"`   // "http://i0.hdslb.com/bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png"
		Theme      struct {
		} `json:"theme"`
		SysNotice struct {
		} `json:"sys_notice"`
		LiveRoom struct {
			RoomStatus    int    `json:"roomStatus"`     // 1
			LiveStatus    int    `json:"liveStatus"`     // 0
			URL           string `json:"url"`            // "https://live.bilibili.com/3394945?broadcast_type=0&is_room_feed=0"
			Title         string `json:"title"`          // "初来报道，关注一下吧~"
			Cover         string `json:"cover"`          // "https://i0.hdslb.com/bfs/live/59fc254c1f51a962dbf69ae85e4920f2f6fb8dcd.png"
			RoomID        int    `json:"roomid"`         // 3394945
			RoundStatus   int    `json:"roundStatus"`    // 0
			BroadcastType int    `json:"broadcast_type"` // 0
			WatchedShow   struct {
				Switch       bool   `json:"switch"`        // true
				Num          int    `json:"num"`           // 1
				TextSmall    string `json:"text_small"`    // "1"
				TextLarge    string `json:"text_large"`    // "1人看过"
				Icon         string `json:"icon"`          // "https://i0.hdslb.com/bfs/live/a725a9e61242ef44d764ac911691a7ce07f36c1d.png"
				IconLocation string `json:"icon_location"` // ""
				IconWeb      string `json:"icon_web"`      // "https://i0.hdslb.com/bfs/live/8d9d0f33ef8bf6f308742752d13dd0df731df19c.png"
			} `json:"watched_show"`
		} `json:"live_room"`
		Birthday string `json:"birthday"` // "01-23"
		School   struct {
			Name string `json:"name"` // "成都信息工程大学"
		} `json:"school"`
		Profession struct {
			Name       string `json:"name"`       // ""
			Department string `json:"department"` // ""
			Title      string `json:"title"`      // ""
			IsShow     int    `json:"is_show"`    // 0
		} `json:"profession"`
		Tags   any `json:"tags"`
		Series struct {
			UserUpgradeStatus int  `json:"user_upgrade_status"` // 3
			ShowUpgradeWindow bool `json:"show_upgrade_window"` // false
		} `json:"series"`
		IsSeniorMember int  `json:"is_senior_member"` // 0
		McnInfo        any  `json:"mcn_info"`
		GaiaResType    int  `json:"gaia_res_type"` // 0
		GaiaData       any  `json:"gaia_data"`
		IsRisk         bool `json:"is_risk"` // false
		Elec           struct {
			ShowInfo struct {
				Show    bool   `json:"show"`     // false
				State   int    `json:"state"`    // -1
				Title   string `json:"title"`    // ""
				Icon    string `json:"icon"`     // ""
				JumpURL string `json:"jump_url"` // "?oid=208259"
			} `json:"show_info"`
		} `json:"elec"`
		Contract struct {
			IsDisplay       bool `json:"is_display"`        // false
			IsFollowDisplay bool `json:"is_follow_display"` // false
		} `json:"contract"`
		CertificateShow bool `json:"certificate_show"` // false
		NameRender      any  `json:"name_render"`
	} `json:"data"`
}

// 用户空间详细信息
func GetAccInfo(uid int, credential *Credential) (result AccInfoResponse, err error) {
	err = cli.Result(AccInfo{MID: uid, Credential: credential}, &result)
	return
}

// 用户名片信息
type Card struct {
	req.Get
	*Credential

	// 目标用户 UID
	MID int `api:"query"`

	// 是否请求用户主页头图
	Photo bool `api:"query"`
}

func (Card) RawURL() string {
	return "/web-interface/card"
}

type CardResponse struct {
	Error
	Data struct {
		Card struct {
			MID         string `json:"mid"`           // "208259"
			Name        string `json:"name"`          // "陈睿"
			Approve     bool   `json:"approve"`       // false
			Sex         string `json:"sex"`           // "男"
			Rank        string `json:"rank"`          // "10000"
			Face        string `json:"face"`          // "http://i2.hdslb.com/bfs/app/8920e6741fc2808cce5b81bc27abdbda291655d3.png"
			FaceNft     int    `json:"face_nft"`      // 0
			FaceNftType int    `json:"face_nft_type"` // 0
			DisplayRank string `json:"DisplayRank"`   // "0"
			Regtime     int    `json:"regtime"`       // 0
			Spacesta    int    `json:"spacesta"`      // 0
			Birthday    string `json:"birthday"`      // ""
			Place       string `json:"place"`         // ""
			Description string `json:"description"`   // ""
			Article     int    `json:"article"`       // 0
			Attentions  []any  `json:"attentions"`
			Fans        int    `json:"fans"`      // 2280857
			Friend      int    `json:"friend"`    // 528
			Attention   int    `json:"attention"` // 528
			Sign        string `json:"sign"`      // "喜欢的话就坚持吧"
			LevelInfo   struct {
				CurrentLevel int `json:"current_level"` // 6
				CurrentMin   int `json:"current_min"`   // 0
				CurrentExp   int `json:"current_exp"`   // 0
				NextExp      int `json:"next_exp"`      // 0
			} `json:"level_info"`
			Pendant struct {
				PID               int    `json:"pid"`                 // 0
				Name              string `json:"name"`                // ""
				Image             string `json:"image"`               // ""
				Expire            int    `json:"expire"`              // 0
				ImageEnhance      string `json:"image_enhance"`       // ""
				ImageEnhanceFrame string `json:"image_enhance_frame"` // ""
				NPID              int    `json:"n_pid"`               // 0
			} `json:"pendant"`
			Nameplate struct {
				NID        int    `json:"nid"`         // 0
				Name       string `json:"name"`        // ""
				Image      string `json:"image"`       // ""
				ImageSmall string `json:"image_small"` // ""
				Level      string `json:"level"`       // ""
				Condition  string `json:"condition"`   // ""
			} `json:"nameplate"`
			Official struct {
				Role  int    `json:"role"`  // 2
				Title string `json:"title"` // "bilibili董事长兼CEO"
				Desc  string `json:"desc"`  // ""
				Type  int    `json:"type"`  // 0
			} `json:"Official"`
			OfficialVerify struct {
				Type int    `json:"type"` // 0
				Desc string `json:"desc"` // "bilibili董事长兼CEO"
			} `json:"official_verify"`
			Vip struct {
				Type       int   `json:"type"`         // 2
				Status     int   `json:"status"`       // 1
				DueDate    int64 `json:"due_date"`     // 2096121600000
				VipPayType int   `json:"vip_pay_type"` // 1
				ThemeType  int   `json:"theme_type"`   // 0
				Label      struct {
					Path                  string `json:"path"`                      // ""
					Text                  string `json:"text"`                      // "十年大会员"
					LabelTheme            string `json:"label_theme"`               // "ten_annual_vip"
					TextColor             string `json:"text_color"`                // "#FFFFFF"
					BgStyle               int    `json:"bg_style"`                  // 1
					BgColor               string `json:"bg_color"`                  // "#FB7299"
					BorderColor           string `json:"border_color"`              // ""
					UseImgLabel           bool   `json:"use_img_label"`             // true
					ImgLabelURIHans       string `json:"img_label_uri_hans"`        // "https://i0.hdslb.com/bfs/activity-plat/static/20220608/e369244d0b14644f5e1a06431e22a4d5/wltavwHAkL.gif"
					ImgLabelURIHant       string `json:"img_label_uri_hant"`        // ""
					ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"` // "https://i0.hdslb.com/bfs/vip/802418ff03911645648b63aa193ba67997b5a0bc.png"
					ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"` // "https://i0.hdslb.com/bfs/activity-plat/static/20220614/e369244d0b14644f5e1a06431e22a4d5/8u7iRTPE7N.png"
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`     // 1
				NicknameColor      string `json:"nickname_color"`       // "#FB7299"
				Role               int    `json:"role"`                 // 7
				AvatarSubscriptURL string `json:"avatar_subscript_url"` // ""
				TvVipStatus        int    `json:"tv_vip_status"`        // 1
				TvVipPayType       int    `json:"tv_vip_pay_type"`      // 0
				TvDueDate          int    `json:"tv_due_date"`          // 1873641600
				AvatarIcon         struct {
					IconType     int `json:"icon_type"` // 1
					IconResource struct {
					} `json:"icon_resource"`
				} `json:"avatar_icon"`
				VipType   int `json:"vipType"`   // 2
				VipStatus int `json:"vipStatus"` // 1
			} `json:"vip"`
			IsSeniorMember int `json:"is_senior_member"` // 0
			NameRender     any `json:"name_render"`
		} `json:"card"`
		Space struct {
			SImg string `json:"s_img"` // "http://i0.hdslb.com/bfs/space/768cc4fd97618cf589d23c2711a1d1a729f42235.png"
			LImg string `json:"l_img"` // "http://i0.hdslb.com/bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png"
		} `json:"space"`
		Following    bool `json:"following"`     // false
		ArchiveCount int  `json:"archive_count"` // 13
		ArticleCount int  `json:"article_count"` // 0
		Follower     int  `json:"follower"`      // 2280857
		LikeNum      int  `json:"like_num"`      // 5862111
	} `json:"data"`
}

// 用户名片信息
func GetCard(uid int, credential *Credential) (result CardResponse, err error) {
	err = cli.Result(Card{MID: uid, Photo: true, Credential: credential}, &result)
	return
}

// 登录用户空间详细信息
type MyInfo struct {
	req.Get
	*Credential
}

func (MyInfo) RawURL() string {
	return "/space/myinfo"
}

type MyInfoResponse struct {
	Error
	Data struct {
		MID            int    `json:"mid"`            // 188888131
		Name           string `json:"name"`           // "脆鲨12138"
		Sex            string `json:"sex"`            // "女"
		Face           string `json:"face"`           // "https://i1.hdslb.com/bfs/face/86faab4844dd2c45870fdafa8f2c9ce7be3e999f.jpg"
		Sign           string `json:"sign"`           // "虚拟艺人团体Vcslive成员，喜欢宅在家打游戏的脆鲨12138！！！"
		Rank           int    `json:"rank"`           // 10000
		Level          int    `json:"level"`          // 6
		Jointime       int    `json:"jointime"`       // 0
		Moral          int    `json:"moral"`          // 70
		Silence        int    `json:"silence"`        // 0
		EmailStatus    int    `json:"email_status"`   // 0
		TelStatus      int    `json:"tel_status"`     // 1
		Identification int    `json:"identification"` // 1
		Vip            struct {
			Type       int   `json:"type"`         // 2
			Status     int   `json:"status"`       // 1
			DueDate    int64 `json:"due_date"`     // 1797264000000
			VipPayType int   `json:"vip_pay_type"` // 1
			ThemeType  int   `json:"theme_type"`   // 0
			Label      struct {
				Path                  string `json:"path"`                      // ""
				Text                  string `json:"text"`                      // "年度大会员"
				LabelTheme            string `json:"label_theme"`               // "annual_vip"
				TextColor             string `json:"text_color"`                // "#FFFFFF"
				BgStyle               int    `json:"bg_style"`                  // 1
				BgColor               string `json:"bg_color"`                  // "#FB7299"
				BorderColor           string `json:"border_color"`              // ""
				UseImgLabel           bool   `json:"use_img_label"`             // true
				ImgLabelURIHans       string `json:"img_label_uri_hans"`        // ""
				ImgLabelURIHant       string `json:"img_label_uri_hant"`        // ""
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"` // "https://i0.hdslb.com/bfs/vip/8d4f8bfc713826a5412a0a27eaaac4d6b9ede1d9.png"
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"` // "https://i0.hdslb.com/bfs/activity-plat/static/20220614/e369244d0b14644f5e1a06431e22a4d5/VEW8fCC0hg.png"
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`     // 1
			NicknameColor      string `json:"nickname_color"`       // "#FB7299"
			Role               int    `json:"role"`                 // 3
			AvatarSubscriptURL string `json:"avatar_subscript_url"` // ""
			TvVipStatus        int    `json:"tv_vip_status"`        // 0
			TvVipPayType       int    `json:"tv_vip_pay_type"`      // 0
			TvDueDate          int    `json:"tv_due_date"`          // 0
			AvatarIcon         struct {
				IconType     int `json:"icon_type"` // 1
				IconResource struct {
				} `json:"icon_resource"`
			} `json:"avatar_icon"`
		} `json:"vip"`
		Pendant struct {
			PID               int    `json:"pid"`                 // 4820
			Name              string `json:"name"`                // "七海演唱会"
			Image             string `json:"image"`               // "https://i1.hdslb.com/bfs/garb/item/7061b098537482bb78242a16cba5d44b426b429d.png"
			Expire            int    `json:"expire"`              // 0
			ImageEnhance      string `json:"image_enhance"`       // "https://i1.hdslb.com/bfs/garb/item/7061b098537482bb78242a16cba5d44b426b429d.png"
			ImageEnhanceFrame string `json:"image_enhance_frame"` // ""
			NPID              int    `json:"n_pid"`               // 4820
		} `json:"pendant"`
		Nameplate struct {
			NID        int    `json:"nid"`         // 62
			Name       string `json:"name"`        // "有爱大佬"
			Image      string `json:"image"`       // "https://i0.hdslb.com/bfs/face/a10ee6b613e0d68d2dfdac8bbf71b94824e10408.png"
			ImageSmall string `json:"image_small"` // "https://i0.hdslb.com/bfs/face/54f4c31ab9b1f1fa2c29dbbc967f66535699337e.png"
			Level      string `json:"level"`       // "普通勋章"
			Condition  string `json:"condition"`   // "当前持有粉丝勋章最高等级>=15级"
		} `json:"nameplate"`
		Official struct {
			Role  int    `json:"role"`  // 0
			Title string `json:"title"` // ""
			Desc  string `json:"desc"`  // ""
			Type  int    `json:"type"`  // -1
		} `json:"official"`
		Birthday      int  `json:"birthday"`
		IsTourist     int  `json:"is_tourist"`      // 0
		IsFakeAccount int  `json:"is_fake_account"` // 0
		PinPrompting  int  `json:"pin_prompting"`   // 0
		IsDeleted     int  `json:"is_deleted"`      // 0
		InRegAudit    int  `json:"in_reg_audit"`    // 0
		IsRipUser     bool `json:"is_rip_user"`     // false
		Profession    struct {
			ID              int    `json:"id"`               // 0
			Name            string `json:"name"`             // ""
			ShowName        string `json:"show_name"`        // ""
			IsShow          int    `json:"is_show"`          // 0
			CategoryOne     string `json:"category_one"`     // ""
			Realname        string `json:"realname"`         // ""
			Title           string `json:"title"`            // ""
			Department      string `json:"department"`       // ""
			CertificateNo   string `json:"certificate_no"`   // ""
			CertificateShow bool   `json:"certificate_show"` // false
		} `json:"profession"`
		FaceNft        int `json:"face_nft"`         // 0
		FaceNftNew     int `json:"face_nft_new"`     // 0
		IsSeniorMember int `json:"is_senior_member"` // 0
		Honours        struct {
			MID    int `json:"mid"` // 188888131
			Colour struct {
				Dark   string `json:"dark"`   // "#CE8620"
				Normal string `json:"normal"` // "#F0900B"
			} `json:"colour"`
			Tags              any `json:"tags"`
			IsLatest100Honour int `json:"is_latest_100honour"` // 0
		} `json:"honours"`
		DigitalID   string `json:"digital_id"`   // ""
		DigitalType int    `json:"digital_type"` // -2
		Attestation struct {
			Type       int `json:"type"` // 0
			CommonInfo struct {
				Title       string `json:"title"`        // ""
				Prefix      string `json:"prefix"`       // ""
				PrefixTitle string `json:"prefix_title"` // ""
			} `json:"common_info"`
			SpliceInfo struct {
				Title string `json:"title"` // ""
			} `json:"splice_info"`
			Icon string `json:"icon"` // ""
			Desc string `json:"desc"` // ""
		} `json:"attestation"`
		ExpertInfo struct {
			Title string `json:"title"` // ""
			State int    `json:"state"` // 0
			Type  int    `json:"type"`  // 0
			Desc  string `json:"desc"`  // ""
		} `json:"expert_info"`
		NameRender  any    `json:"name_render"`
		CountryCode string `json:"country_code"` // "86"
		LevelExp    struct {
			CurrentLevel int   `json:"current_level"` // 6
			CurrentMin   int   `json:"current_min"`   // 28800
			CurrentExp   int   `json:"current_exp"`   // 48539
			NextExp      int   `json:"next_exp"`      // -1
			LevelUp      int64 `json:"level_up"`      // -62135596800
		} `json:"level_exp"`
		Coins     float64 `json:"coins"`     // 1547.8
		Following int     `json:"following"` // 6601
		Follower  int     `json:"follower"`  // 6601
	} `json:"data"`
}

// 登录用户空间详细信息
func GetMyInfo(credential *Credential) (result MyInfoResponse, err error) {
	err = cli.Result(MyInfo{Credential: credential}, &result)
	return
}

// 多用户详细信息字典
type CardMap struct {
	req.Get
	*Credential

	// 目标用户的 UID 列表 用(,)间隔
	// 最多200个成员
	UIDs string `api:"query"`
}

func (CardMap) RawURL() string {
	return "/polymer/pc-electron/v1/user/cards"
}

type CardMapResponse struct {
	Error
	Data map[string]struct {
		Face       string `json:"face"`         // "https://i1.hdslb.com/bfs/face/8802d20c0c0e7f4cc232543c619da14ec5bac76d.jpg"
		FaceNft    int    `json:"face_nft"`     // 0
		FaceNftNew int    `json:"face_nft_new"` // 0
		MID        string `json:"mid"`          // "434324701"
		Name       string `json:"name"`         // "奔跑在上岸路上的小杨"
		NameRender any    `json:"name_render"`
		Nameplate  any    `json:"nameplate"`
		Official   struct {
			Desc  string `json:"desc"`  // ""
			Role  int    `json:"role"`  // 0
			Title string `json:"title"` // ""
			Type  int    `json:"type"`  // -1
		} `json:"official"`
		Pendant any `json:"pendant"`
		Vip     struct {
			AvatarIcon         any    `json:"avatar_icon"`
			AvatarSubscript    int    `json:"avatar_subscript"`     // 0
			AvatarSubscriptURL string `json:"avatar_subscript_url"` // ""
			DueDate            string `json:"due_date"`             // "0"
			Label              struct {
				BgColor               string `json:"bg_color"`                  // ""
				BgStyle               int    `json:"bg_style"`                  // 0
				BorderColor           string `json:"border_color"`              // ""
				ImgLabelURIHans       string `json:"img_label_uri_hans"`        // ""
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"` // "https://i0.hdslb.com/bfs/vip/d7b702ef65a976b20ed854cbd04cb9e27341bb79.png"
				ImgLabelURIHant       string `json:"img_label_uri_hant"`        // ""
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"` // "https://i0.hdslb.com/bfs/activity-plat/static/20220614/e369244d0b14644f5e1a06431e22a4d5/KJunwh19T5.png"
				LabelTheme            string `json:"label_theme"`               // ""
				Path                  string `json:"path"`                      // ""
				Text                  string `json:"text"`                      // ""
				TextColor             string `json:"text_color"`                // ""
				UseImgLabel           bool   `json:"use_img_label"`             // true
			} `json:"label"`
			NicknameColor string `json:"nickname_color"`  // ""
			Role          string `json:"role"`            // "0"
			Status        int    `json:"status"`          // 0
			ThemeType     int    `json:"theme_type"`      // 0
			TvDueDate     string `json:"tv_due_date"`     // "0"
			TvVipPayType  int    `json:"tv_vip_pay_type"` // 0
			TvVipStatus   int    `json:"tv_vip_status"`   // 0
			Type          int    `json:"type"`            // 0
			VipPayType    int    `json:"vip_pay_type"`    // 0
		} `json:"vip"`
	} `json:"data"`
}

// 多用户详细信息字典
func GetCardMap(uid []int, credential *Credential) (result CardMapResponse, err error) {
	err = cli.Result(CardMap{UIDs: IntSliceToString(uid), Credential: credential}, &result)
	return
}

// 多用户详细信息切片
type CardSlice struct {
	req.Get
	*Credential

	// 目标用户的 UID 列表 用(,)间隔
	// 最多 50 个成员
	UIDs string `api:"query"`
}

func (CardSlice) RawURL() string {
	return "https://api.vc.bilibili.com/account/v1/user/cards"
}

type CardSliceResponse struct {
	Error
	Data []struct {
		MID     int    `json:"mid"`     // 7706705
		Name    string `json:"name"`    // "阿梓从小就很可爱"
		Face    string `json:"face"`    // "https://i2.hdslb.com/bfs/face/217b2cd3d069325244f7f64ff3d75b33c71f43c8.jpg"
		Sign    string `json:"sign"`    // "虚拟艺人团体VirtuaReal成员，喜欢宅在家唱歌打游戏的梓宝！ 商务合作：bd@vrp.live"
		Rank    int    `json:"rank"`    // 10000
		Level   int    `json:"level"`   // 6
		Silence int    `json:"silence"` // 0
	} `json:"data"`
}

// 多用户详细信息切片
func GetCardSlice(uid []int, credential *Credential) (result CardSliceResponse, err error) {
	err = cli.Result(CardSlice{UIDs: IntSliceToString(uid), Credential: credential}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/status_number.html

// 关系状态数
type RelationStat struct {
	req.Get
	*Credential

	// 目标用户 UID
	Vmid int `api:"query"`
}

func (RelationStat) RawURL() string {
	return "/relation/stat"
}

type RelationStatResponse struct {
	Error
	Data struct {
		MID       int `json:"mid"`       // 434334701
		Following int `json:"following"` // 592
		Whisper   int `json:"whisper"`   // 0
		Black     int `json:"black"`     // 0
		Follower  int `json:"follower"`  // 1057173
	} `json:"data"`
}

// 关系状态数
func GetRelationStat(uid int, credential *Credential) (result RelationStatResponse, err error) {
	err = cli.Result(RelationStat{Vmid: uid, Credential: credential}, &result)
	return
}

// UP 主状态数
type UPStat struct {
	req.Get
	*Credential

	// 目标用户 UID
	MID int `api:"query"`
}

func (UPStat) RawURL() string {
	return "/space/upstat"
}

type UPStatResponse struct {
	Error
	Data struct {
		Archive struct {
			EnableVt int `json:"enable_vt"` // 0
			View     int `json:"view"`      // 57937010
			Vt       int `json:"vt"`        // 0
		} `json:"archive"`
		Article struct {
			View int `json:"view"` // 123793
		} `json:"article"`
		Likes int `json:"likes"` // 9789775
	} `json:"data"`
}

// UP 主状态数
func GetUPStat(uid int, credential *Credential) (result UPStatResponse, err error) {
	err = cli.Result(UPStat{MID: uid, Credential: credential}, &result)
	return
}

// 用户导航栏状态数
//
// 但是不知道为什么被禁止访问了
//
// 错误号:412
// 由于触发哔哩哔哩安全风控策略，该次访问请求被拒绝。
// The request was rejected because of the bilibili security control policy.
type NavNum struct {
	req.Get

	// 目标用户 UID
	MID int `api:"query"`
}

func (NavNum) RawURL() string {
	return "/space/navnum"
}

// 相簿投稿数
type UploadCount struct {
	req.Get

	// 目标用户 UID
	UID int `api:"query"`
}

func (UploadCount) RawURL() string {
	return "https://api.vc.bilibili.com/link_draw/v1/doc/upload_count"
}

type UploadCountResponse struct {
	Error
	Data struct {
		AllCount   int `json:"all_count"`   // 1018
		DrawCount  int `json:"draw_count"`  // 0
		PhotoCount int `json:"photo_count"` // 0
		DailyCount int `json:"daily_count"` // 1018
	} `json:"data"`
}

// 相簿投稿数
func GetUploadCount(uid int) (result UploadCountResponse, err error) {
	err = cli.Result(UploadCount{UID: uid}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/relation.html

// 契约计划相关信息
type ContractInfo struct {
	// 目标用户是否为对方的契约者
	IsContract bool `json:"is_contract"`

	// 是否为目标用户的契约者
	IsContractor bool `json:"is_contractor"`

	// 对方成为目标用户的契约者的时间 秒级时间戳
	// 仅当 IsContractor 项的值为 true 时才有此项
	TS int `json:"ts"`

	// 对方作为目标用户的契约者的属性
	// (1)老粉
	// 否则为原始粉丝
	UserAttr int `json:"user_attr"`
}

// 认证信息
type OfficialVerify struct {
	// 用户认证类型
	// (-1)无
	// (0)UP 主认证
	// (1)机构认证
	Type int `json:"type"`

	// 用户认证信息
	Desc string `json:"desc"`
}

// 会员信息
type VIP struct {
	// 会员类型
	// (0)无
	// (1)月度大会员
	// (2)年度以上大会员
	VIPType int `json:"vipType"`

	// 会员到期时间	毫秒级时间戳
	VIPDueDate int64 `json:"vipDueDate"`

	DueRemark    string `json:"dueRemark"`
	AccessStatus int    `json:"accessStatus"`

	// 大会员状态
	// (0)无
	// (1)有
	VIPStatus int `json:"vipStatus"`

	VIPStatusWarn string `json:"vipStatusWarn"`
	ThemeType     int    `json:"themeType"`
	Label         struct {
		Path        string `json:"path"`
		Text        string `json:"text"`
		LabelTheme  string `json:"label_theme"`
		TextColor   string `json:"text_color"`
		BgStyle     int    `json:"bg_style"`
		BgColor     string `json:"bg_color"`
		BorderColor string `json:"border_color"`
	} `json:"label"`
	AvatarSubscript    int    `json:"avatar_subscript"`
	NicknameColor      string `json:"nickname_color"`
	AvatarSubscriptURL string `json:"avatar_subscript_url"`
}

// 直播信息
type Live struct {
	// 开播状态
	// (0)未开播
	// (1)已开播
	LiveStatus int `json:"live_status"`

	// 直播链接
	JumpURL string `json:"jump_url"`
}

// 关系列表对象
type RelationItem struct {
	// 用户 UID
	MID int64 `json:"mid"`

	// 对方对于自己的关系属性
	// (0)未关注
	// (1)悄悄关注(现已下线)
	// (2)已关注
	// (6)已互粉
	// (128)已拉黑
	Attribute int `json:"attribute"`

	// 对方关注目标用户时间
	// 秒级时间戳 互关后刷新
	Mtime int `json:"mtime"`

	// 目标用户将对方分组到的 id
	Tag []int `json:"tag"`

	// 特别关注标志
	// (0)否
	// (1)是
	Special int `json:"special"`

	// 契约计划相关信息
	ContractInfo ContractInfo `json:"contract_info"`

	// 用户昵称
	Uname string `json:"uname"`

	// 用户头像 URL
	Face string `json:"face"`

	// 用户签名
	Sign string `json:"sign"`

	// 是否为 NFT 头像
	// (0)非 NFT 头像
	// (1)是 NFT 头像
	FaceNFT int `json:"face_nft"`

	// 认证信息
	OfficialVerify OfficialVerify `json:"official_verify"`

	// 会员信息
	VIP VIP `json:"vip"`

	// 直播信息
	Live Live `json:"live"`

	NameRender struct {
	} `json:"name_render"`
	NFTIcon    string `json:"nft_icon"`
	RecReason  string `json:"rec_reason"`
	TrackID    string `json:"track_id"`
	FollowTime string `json:"follow_time"`
}

func (r RelationItem) String() string {
	return r.Uname
}

// 查询用户粉丝明细
type Followers struct {
	req.Get
	*Credential

	// 目标用户 UID
	Vmid int `api:"query"`

	// 分页页数 仅可查看前 1000 名粉丝
	Pn int `api:"query:1"`

	// 分页大小
	Ps int `api:"query:50"`
}

func (Followers) RawURL() string {
	return "/relation/followers"
}

func (api *Followers) ReadPage() (v FollowersResponse, err error) {
	if api.Pn == 0 {
		api.Pn = 1
	}
	err = cli.Result(api, &v)
	if err != nil {
		return
	}
	if len(v.Data.List) == 0 {
		err = ErrNoMorePage
		return
	}
	api.Pn++
	return
}

var _ PageReader[FollowersResponse] = (*Followers)(nil)

type FollowersResponse struct {
	Error
	Data struct {
		List      []RelationItem `json:"list"`
		ReVersion int            `json:"re_version"` // 0
		Total     int            `json:"total"`      // 1057249
	} `json:"data"`
}

// 查询用户粉丝明细
func GetFollowers(uid int, credential *Credential) (result FollowersResponse, err error) {
	err = cli.Result(Followers{Vmid: uid, Credential: credential}, &result)
	return
}

// 查询用户关注明细
type RelationFollowings struct {
	req.Get
	*Credential

	// 目标用户 UID
	Vmid int `api:"query"`

	// 排序方式 当目标用户为自己时有效
	// ("")按照关注顺序排列
	// ("attention")按照最常访问排列
	OrderType string `api:"query,omitempty"`

	// 分页页数
	Pn int `api:"query:1"`

	// 分页大小
	Ps int `api:"query:50"`
}

func (RelationFollowings) RawURL() string {
	return "/relation/followings"
}

func (api *RelationFollowings) ReadPage() (v RelationFollowingsResponse, err error) {
	if api.Pn == 0 {
		api.Pn = 1
	}
	err = cli.Result(api, &v)
	if err != nil {
		return
	}
	if len(v.Data.List) == 0 {
		err = ErrNoMorePage
		return
	}
	api.Pn++
	return
}

var _ PageReader[RelationFollowingsResponse] = (*RelationFollowings)(nil)

type RelationFollowingsResponse struct {
	Error
	Data struct {
		List      []RelationItem `json:"list"`
		ReVersion int            `json:"re_version"` // 0
		Total     int            `json:"total"`      // 207
	} `json:"data"`
}

// 查询用户关注明细
func GetRelationFollowings(uid int, credential *Credential) (result RelationFollowingsResponse, err error) {
	err = cli.Result(RelationFollowings{Vmid: uid, Credential: credential}, &result)
	return
}

// 搜索关注明细
type FollowingsSearch struct {
	req.Get
	*Credential

	// 目标用户 UID
	Vmid int `api:"query"`

	// 搜索关键词
	Name string `api:"query"`

	// 分页页数
	Pn int `api:"query:1"`

	// 分页大小
	Ps int `api:"query:50"`
}

func (FollowingsSearch) RawURL() string {
	return "/relation/followings/search"
}

func (api *FollowingsSearch) ReadPage() (v FollowingsSearchResponse, err error) {
	if api.Pn == 0 {
		api.Pn = 1
	}
	err = cli.Result(api, &v)
	if err != nil {
		return
	}
	if len(v.Data.List) == 0 {
		err = ErrNoMorePage
		return
	}
	api.Pn++
	return
}

var _ PageReader[FollowingsSearchResponse] = (*FollowingsSearch)(nil)

type FollowingsSearchResponse struct {
	Error
	Data struct {
		List  []RelationItem `json:"list"`
		Total int            `json:"total"` // 1
	} `json:"data"`
}

// 搜索关注明细
func GetFollowingsSearch(uid int, name string, credential *Credential) (result FollowingsSearchResponse, err error) {
	err = cli.Result(FollowingsSearch{Vmid: uid, Name: name, Credential: credential}, &result)
	return
}

// 查询共同关注明细
type SameFollowings struct {
	req.Get
	*Credential

	// 目标用户 UID
	Vmid int `api:"query"`

	// 分页页数
	Pn int `api:"query:1"`

	// 分页大小
	Ps int `api:"query:50"`
}

func (SameFollowings) RawURL() string {
	return "/relation/same/followings"
}

func (api *SameFollowings) ReadPage() (v SameFollowingsResponse, err error) {
	if api.Pn == 0 {
		api.Pn = 1
	}
	err = cli.Result(api, &v)
	if err != nil {
		return
	}
	if len(v.Data.List) == 0 {
		err = ErrNoMorePage
		return
	}
	api.Pn++
	return
}

var _ PageReader[SameFollowingsResponse] = (*SameFollowings)(nil)

type SameFollowingsResponse struct {
	Error
	Data struct {
		List      []RelationItem `json:"list"`
		ReVersion string         `json:"re_version"` // ""
		Total     int            `json:"total"`      // 7
	} `json:"data"`
}

// 查询共同关注明细
func GetSameFollowings(uid int, credential *Credential) (result SameFollowingsResponse, err error) {
	err = cli.Result(SameFollowings{Vmid: uid, Credential: credential}, &result)
	return
}

// 查询悄悄关注明细
type Whispers struct {
	req.Get
	*Credential
}

func (Whispers) RawURL() string {
	return "/relation/whispers"
}

type WhispersResponse struct {
	Error
	Data struct {
		List      []RelationItem `json:"list"`
		ReVersion int            `json:"re_version"` // 0
	} `json:"data"`
}

// 查询悄悄关注明细
func GetWhispers(credential *Credential) (result WhispersResponse, err error) {
	err = cli.Result(Whispers{Credential: credential}, &result)
	return
}

// 查询互相关注明细
type Friends struct {
	req.Get
	*Credential
}

func (Friends) RawURL() string {
	return "/relation/friends"
}

type FriendsResponse struct {
	Error
	Data struct {
		List      []RelationItem `json:"list"`
		ReVersion int            `json:"re_version"` // 0
	} `json:"data"`
}

// 查询互相关注明细
func GetFriends(credential *Credential) (result FriendsResponse, err error) {
	err = cli.Result(Friends{Credential: credential}, &result)
	return
}

// 查询黑名单明细
type Blacks struct {
	req.Get
	*Credential

	// 分页页数
	Pn int `api:"query:1"`

	// 分页大小
	Ps int `api:"query:50"`
}

func (Blacks) RawURL() string {
	return "/relation/blacks"
}

func (api *Blacks) ReadPage() (v BlacksResponse, err error) {
	if api.Pn == 0 {
		api.Pn = 1
	}
	err = cli.Result(api, &v)
	if err != nil {
		return
	}
	if len(v.Data.List) == 0 {
		err = ErrNoMorePage
		return
	}
	api.Pn++
	return
}

var _ PageReader[BlacksResponse] = (*Blacks)(nil)

type BlacksResponse struct {
	Error
	Data struct {
		List      []RelationItem `json:"list"`
		ReVersion int            `json:"re_version"` // 0
		Total     int            `json:"total"`      // 87
	} `json:"data"`
}

// 查询黑名单明细
func GetBlacks(credential *Credential) (result BlacksResponse, err error) {
	err = cli.Result(Blacks{Credential: credential}, &result)
	return
}

// 操作用户关系
type RelationModify struct {
	PostCSRF
	*Credential
	// TODO
}

func (RelationModify) RawURL() string {
	return "/relation/modify"
}

// 批量操作用户关系
type BatchModify struct {
	PostCSRF
	*Credential
	// TODO
}

func (BatchModify) RawURL() string {
	return "/relation/batch/modify"
}

// 查询关系属性
type QueryRelation struct {
	// 目标用户 UID
	MID int `json:"mid"`

	// 关系属性
	// (0)未关注
	// (2)已关注
	// (6)已互粉
	// (128)已拉黑
	Attribute int `json:"attribute"`

	// 关注对方时间 秒级时间戳
	Mtime int `json:"mtime"`

	// 分组 id
	Tag []int `json:"tag"`

	// 特别关注标志
	// (0)否
	// (1)是
	Special int `json:"special"`
}

// 查询用户与自己关系（仅关注）
type Relation struct {
	req.Get
	*Credential

	// 目标用户 UID
	FID int `api:"query"`
}

func (Relation) RawURL() string {
	return "/relation"
}

type RelationResponse struct {
	Error
	Data QueryRelation `json:"data"`
}

// 查询用户与自己关系（仅关注）
func GetRelation(uid int, credential *Credential) (result RelationResponse, err error) {
	err = cli.Result(Relation{FID: uid, Credential: credential}, &result)
	return
}

// 查询用户与自己关系（互相关系）
type AccRelation struct {
	GetWBI
	*Credential

	// 目标用户 UID
	MID int `api:"query"`
}

func (AccRelation) RawURL() string {
	return "/space/wbi/acc/relation"
}

type AccRelationResponse struct {
	Error
	Data struct {
		Relation   QueryRelation `json:"relation"`    // 目标用户对于当前用户的关系
		BeRelation QueryRelation `json:"be_relation"` // 当前用户对于目标用户的关系
	} `json:"data"`
}

// 查询用户与自己关系（互相关系）
func GetAccRelation(uid int, credential *Credential) (result AccRelationResponse, err error) {
	err = cli.Result(AccRelation{MID: uid, Credential: credential}, &result)
	return
}

// 批量查询用户与自己关系
type Relations struct {
	req.Get
	*Credential

	// 目标用户 UID 用(,)间隔
	FIDs string `api:"query"`
}

func (Relations) RawURL() string {
	return "/relation/relations"
}

type RelationsResponse struct {
	Error
	Data map[string]QueryRelation `json:"data"`
}

// 批量查询用户与自己关系
func GetRelations(uid []int, credential *Credential) (result RelationsResponse, err error) {
	err = cli.Result(Relations{FIDs: IntSliceToString(uid), Credential: credential}, &result)
	return
}

// 查询关注分组列表
type Tags struct {
	req.Get
	*Credential
}

func (Tags) RawURL() string {
	return "/relation/tags"
}

type TagsResponse struct {
	Error
	Data []struct {
		TagID int    `json:"tagid"` // -10
		Name  string `json:"name"`  // "特别关注"
		Count int    `json:"count"` // 2
		Tip   string `json:"tip"`   // "第一时间收到该分组下用户更新稿件的通知"
	} `json:"data"`
}

// 查询关注分组列表
func GetTags(credential *Credential) (result TagsResponse, err error) {
	err = cli.Result(Tags{Credential: credential}, &result)
	return
}

// 查询关注分组明细
type Tag struct {
	req.Get
	*Credential

	// 分组 id
	TagID int `api:"query" req:"tagid"`

	// 排序方式 当目标用户为自己时有效
	// ("")按照关注顺序排列
	// ("attention")按照最常访问排列
	OrderType string `api:"query,omitempty"`

	// 分页页数
	Pn int `api:"query:1"`

	// 分页大小
	Ps int `api:"query:20"`
}

func (Tag) RawURL() string {
	return "/relation/tag"
}

func (api *Tag) ReadPage() (v TagResponse, err error) {
	if api.Pn == 0 {
		api.Pn = 1
	}
	err = cli.Result(api, &v)
	if err != nil {
		return
	}
	if len(v.Data) == 0 {
		err = ErrNoMorePage
		return
	}
	api.Pn++
	return
}

var _ PageReader[TagResponse] = (*Tag)(nil)

type TagResponse struct {
	Error
	Data []RelationItem `json:"data"`
}

// 查询关注分组明细
func GetTag(tagid int, credential *Credential) (result TagResponse, err error) {
	err = cli.Result(Tag{TagID: tagid, Credential: credential}, &result)
	return
}

// 查询目标用户所在的分组
type User struct {
	req.Get
	*Credential

	// 目标用户 UID
	FID int `api:"query"`
}

func (User) RawURL() string {
	return "/relation/tag/user"
}

type UserResponse struct {
	Error
	Data map[string]string `json:"data"`
}

// 查询目标用户所在的分组
func GetUser(uid int, credential *Credential) (result UserResponse, err error) {
	err = cli.Result(User{FID: uid, Credential: credential}, &result)
	return
}

// 查询所有特别关注 UID
type Special struct {
	req.Get
	*Credential
}

func (Special) RawURL() string {
	return "/relation/tag/special"
}

type SpecialResponse struct {
	Error
	Data []int `json:"data"`
}

// 查询所有特别关注 UID
func GetSpecial(credential *Credential) (result SpecialResponse, err error) {
	err = cli.Result(Special{Credential: credential}, &result)
	return
}

// 创建分组
type Create struct {
	PostCSRF
	*Credential

	// 分组名 最长 16 字符
	Tag string `api:"body"`
}

func (Create) RawURL() string {
	return "/relation/tag/create"
}

type CreateResponse struct {
	Error
	Data struct {
		TagID int `json:"tagid"`
	} `json:"data"`
}

// 创建分组
func PostCreate(tag string, credential *Credential) (result CreateResponse, err error) {
	err = cli.Result(Create{Tag: tag, Credential: credential}, &result)
	return
}

// 重命名分组
type Update struct {
	PostCSRF
	*Credential

	// 分组 id
	TagID int `api:"body" req:"tagid"`

	// 新分组名 最长 16 字符
	Name string `api:"body"`
}

func (Update) RawURL() string {
	return "/relation/tag/update"
}

type UpdateResponse struct {
	Error
}

// 重命名分组
func PostUpdate(tag int, name string, credential *Credential) (result UpdateResponse, err error) {
	err = cli.Result(Update{TagID: tag, Name: name, Credential: credential}, &result)
	return
}

// 删除分组
type Del struct {
	PostCSRF
	*Credential

	// 分组 id
	TagID int `api:"body" req:"tagid"`
}

func (Del) RawURL() string {
	return "/relation/tag/del"
}

type DelResponse struct {
	Error
}

// 删除分组
func PostDel(tagid int, credential *Credential) (result DelResponse, err error) {
	err = cli.Result(Del{TagID: tagid, Credential: credential}, &result)
	return
}

// 修改分组成员
//
// 如需删除分组中的成员 请将 tagids 设为 0 即移动至默认分组 而不是取关
type AddUsers struct {
	PostCSRF
	*Credential

	// 目标用户 UID 用(,)间隔
	FIDs string `api:"query"`

	// 分组 id 列表 用(,)间隔
	TagIDs string `api:"query" req:"tagids"`
}

func (AddUsers) RawURL() string {
	return "/relation/tags/addUsers"
}

type AddUsersResponse struct {
	Error
}

// 将数组转为字符串
func IntSliceToString(num []int) string {
	s := make([]string, 0, len(num))
	for _, i := range num {
		s = append(s, strconv.Itoa(i))
	}
	return strings.Join(s, ",")
}

// 修改分组成员
//
// 如需删除分组中的成员 请将 tagids 设为 0 即移动至默认分组 而不是取关
func PostAddUsers(uid, tagid []int, credential *Credential) (result AddUsersResponse, err error) {
	err = cli.Result(AddUsers{FIDs: IntSliceToString(uid), TagIDs: IntSliceToString(tagid), Credential: credential}, &result)
	return
}

// 复制关注到分组
type CopyUsers struct {
	PostCSRF
	*Credential

	// 待复制用户 UID 用(,)间隔
	FIDs string `api:"query"`

	// 目标分组 id 列表 用(,)间隔
	TagIDs string `api:"query" req:"tagids"`
}

func (CopyUsers) RawURL() string {
	return "/relation/tags/copyUsers"
}

type CopyUsersResponse struct {
	Error
}

// 复制关注到分组
func PostCopyUsers(uid, tagid []int, credential *Credential) (result CopyUsersResponse, err error) {
	err = cli.Result(CopyUsers{FIDs: IntSliceToString(uid), TagIDs: IntSliceToString(tagid), Credential: credential}, &result)
	return
}

// 移动关注到分组
type MoveUsers struct {
	PostCSRF
	*Credential

	// 原分组 id 列表 用(,)间隔
	BeforeTagIDs string `api:"body" req:"beforeTagids"`

	// 新分组 id 列表 用(,)间隔
	AfterTagIDs string `api:"body" req:"afterTagids"`

	// 待移动用户 UID 用(,)间隔
	FIDs string `api:"query"`
}

func (MoveUsers) RawURL() string {
	return "/relation/tags/moveUsers"
}

type MoveUsersResponse struct {
	Error
}

// 移动关注到分组
func PostMoveUsers(uid, beforeTagid, afterTagid []int, credential *Credential) (result MoveUsersResponse, err error) {
	err = cli.Result(MoveUsers{
		BeforeTagIDs: IntSliceToString(beforeTagid),
		AfterTagIDs:  IntSliceToString(afterTagid),
		FIDs:         IntSliceToString(uid),
		Credential:   credential,
	}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/medals.html

// 指定用户的所有粉丝勋章信息
type MedalWall struct {
	req.Get
	*Credential

	// 目标 UID
	TargetID int `api:"query"`
}

func (MedalWall) RawURL() string {
	return "https://api.live.bilibili.com/xlive/web-ucenter/user/MedalWall"
}

type MedalWallResponse struct {
	Error
	Data struct {
		List []struct {
			MedalInfo struct {
				TargetID         int    `json:"target_id"`          // 7706705
				Level            int    `json:"level"`              // 28
				MedalName        string `json:"medal_name"`         // "小孩梓"
				MedalColorStart  int    `json:"medal_color_start"`  // 398668
				MedalColorEnd    int    `json:"medal_color_end"`    // 6850801
				MedalColorBorder int    `json:"medal_color_border"` // 6809855
				GuardLevel       int    `json:"guard_level"`        // 3
				WearingStatus    int    `json:"wearing_status"`     // 1
				MedalID          int    `json:"medal_id"`           // 13139
				Intimacy         int    `json:"intimacy"`           // 66132
				NextIntimacy     int    `json:"next_intimacy"`      // 160000
				TodayFeed        int    `json:"today_feed"`         // 0
				DayLimit         int    `json:"day_limit"`          // 250000
				GuardIcon        string `json:"guard_icon"`         // "https://i0.hdslb.com/bfs/live/143f5ec3003b4080d1b5f817a9efdca46d631945.png"
				HonorIcon        string `json:"honor_icon"`         // ""
			} `json:"medal_info"`
			TargetName string `json:"target_name"` // "阿梓从小就很可爱"
			TargetIcon string `json:"target_icon"` // "https://i2.hdslb.com/bfs/face/217b2cd3d069325244f7f64ff3d75b33c71f43c8.jpg"
			Link       string `json:"link"`        // "https://live.bilibili.com/80397?accept_quality=%5B10000%2C400%2C250%5D&av1_current_qn=0&broadcast_type=0&current_qn=10000&current_quality=10000&h264_current_qn=10000&h265_current_qn=10000&is_room_feed=1&live_from=28013&live_play_network=other&master_url=&p2p_type=-1&playurl_av1=&playurl_h264=http%3A%2F%2F123.234.3.207%2Flive-bvc%2F590881%2Flive_7706705_4063361_bluray.flv%3Fexpires%3D1735143025%26pt%3Dandroid%26deadline%3D1735143025%26len%3D0%26oi%3D1869969956%26platform%3Dandroid%26qn%3D10000%26trid%3D10000cc48f6d1b916d7f07e4c0e4be676c20%26uipk%3D100%26uipv%3D100%26nbs%3D1%26uparams%3Dcdn%2Cdeadline%2Clen%2Coi%2Cplatform%2Cqn%2Ctrid%2Cuipk%2Cuipv%2Cnbs%26cdn%3Dcn-gotcha01%26upsig%3D6e3a89585667751f71d1158bd00b653d%26site%3Df9ee5f4915d5bded29049b85ce6c4bc1%26free_type%3D0%26mid%3D188888131%26sche%3Dban%26sid%3Dcn-sdqd-cu-01-38%26chash%3D1%26bmt%3D1%26sg%3Ddf%26trace%3D4877%26isp%3Dother%26rg%3DNorthEast%26pv%3DLiaoning%26suffix%3Dbluray%26qp%3Dhv_10000%26source%3Dpuv3_batch%26sk%3D1304f646dfeb4df8b6e7ff33c167d3ad491555cae985818bb2a5ec341f048d3d%26deploy_env%3Dprod%26score%3D1%26pp%3Drtmp%26info_source%3Dcache%26hot_cdn%3D909513%26sl%3D4%26p2p_type%3D-1%26origin_bitrate%3D945259%26vd%3Dnc%26zoneid_l%3D151420932%26sid_l%3Dstream_name_cold%26src%3Dpuv3%26order%3D1&playurl_h265=http%3A%2F%2F223.111.250.61%2Flive-bvc%2F534453%2Flive_7706705_4063361_prohevc.flv%3Fexpires%3D1735143025%26pt%3Dandroid%26deadline%3D1735143025%26len%3D0%26oi%3D1869969956%26platform%3Dandroid%26qn%3D10000%26trid%3D10000cc48f6d1b916d7f07e4c0e4be676c20%26uipk%3D100%26uipv%3D100%26nbs%3D1%26uparams%3Dcdn%2Cdeadline%2Clen%2Coi%2Cplatform%2Cqn%2Ctrid%2Cuipk%2Cuipv%2Cnbs%26cdn%3Dcn-gotcha01%26upsig%3Ddf3e1ac3550756b70d49fb6468663495%26site%3Df9ee5f4915d5bded29049b85ce6c4bc1%26free_type%3D0%26mid%3D188888131%26sche%3Dban%26sid%3Dcn-jssz-cm-02-76%26chash%3D1%26bmt%3D1%26sg%3Ddf%26trace%3D13%26isp%3Dother%26rg%3DNorthEast%26pv%3DLiaoning%26pp%3Drtmp%26info_source%3Dcache%26suffix%3Dprohevc%26score%3D57%26deploy_env%3Dprod%26origin_bitrate%3D945259%26p2p_type%3D-1%26qp%3Dhv_10000%26source%3Dpuv3_batch%26sk%3D1304f646dfeb4df8b6e7ff33c167d3ad491555cae985818bb2a5ec341f048d3d%26hot_cdn%3D909513%26sl%3D4%26vd%3Dnc%26zoneid_l%3D151420932%26sid_l%3Dlive_7706705_4063361_prohevc%26src%3Dpuv3%26order%3D1&quality_description=%5B%7B%22qn%22%3A10000%2C%22desc%22%3A%22%E5%8E%9F%E7%94%BB%22%2C%22hdr_type%22%3A0%7D%2C%7B%22qn%22%3A400%2C%22desc%22%3A%22%E8%93%9D%E5%85%89%22%2C%22hdr_type%22%3A0%7D%2C%7B%22qn%22%3A250%2C%22desc%22%3A%22%E8%B6%85%E6%B8%85%22%2C%22hdr_type%22%3A0%7D%5D"
			LiveStatus int    `json:"live_status"` // 1
			Official   int    `json:"official"`    // 1
			UinfoMedal struct {
				Name               string `json:"name"`                  // "小孩梓"
				Level              int    `json:"level"`                 // 28
				ColorStart         int    `json:"color_start"`           // 398668
				ColorEnd           int    `json:"color_end"`             // 6850801
				ColorBorder        int    `json:"color_border"`          // 6809855
				Color              int    `json:"color"`                 // 0
				ID                 int    `json:"id"`                    // 13139
				Typ                int    `json:"typ"`                   // 0
				IsLight            int    `json:"is_light"`              // 1
				RUID               int    `json:"ruid"`                  // 7706705
				GuardLevel         int    `json:"guard_level"`           // 3
				Score              int    `json:"score"`                 // 0
				GuardIcon          string `json:"guard_icon"`            // "https://i0.hdslb.com/bfs/live/143f5ec3003b4080d1b5f817a9efdca46d631945.png"
				HonorIcon          string `json:"honor_icon"`            // ""
				V2MedalColorStart  string `json:"v2_medal_color_start"`  // "#4775EFCC"
				V2MedalColorEnd    string `json:"v2_medal_color_end"`    // "#4775EFCC"
				V2MedalColorBorder string `json:"v2_medal_color_border"` // "#58A1F8FF"
				V2MedalColorText   string `json:"v2_medal_color_text"`   // "#FFFFFFFF"
				V2MedalColorLevel  string `json:"v2_medal_color_level"`  // "#000B7099"
				UserReceiveCount   int    `json:"user_receive_count"`    // 0
			} `json:"uinfo_medal"`
		} `json:"list"`
		Count           int    `json:"count"`             // 1
		CloseSpaceMedal int    `json:"close_space_medal"` // 0
		OnlyShowWearing int    `json:"only_show_wearing"` // 1
		Name            string `json:"name"`              // "七海Nana7mi"
		Icon            string `json:"icon"`              // "https://i2.hdslb.com/bfs/face/3adb26401cfab0fe6b1a0d5b2c09220499108d64.jpg"
		UID             int    `json:"uid"`               // 188888131 (编者著：神金啊怎么返回的是查询者的 UID)
		Level           int    `json:"level"`             // 6
	} `json:"data"`
}

// 指定用户的所有粉丝勋章信息
func GetMedalWall(uid int, credential *Credential) (result MedalWallResponse, err error) {
	err = cli.Result(MedalWall{TargetID: uid, Credential: credential}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/space.html

// 视频信息
type VideoInfo struct {
	AID       int    `json:"aid"`                  // 78090377
	Videos    int    `json:"videos"`               // 1
	TID       int    `json:"tid"`                  // 47
	Tname     string `json:"tname"`                // "同人·手书"
	Copyright int    `json:"copyright"`            // 1
	Pic       string `json:"pic"`                  // "http://i2.hdslb.com/bfs/archive/7fe8272ef4c90d07ba2dba968638392f8d5bf490.jpg"
	Title     string `json:"title"`                // "【七海】七海的偶像宣言／私、アイドル宣言【手书PV】"
	Pubdate   int    `json:"pubdate"`              // 1575540036
	Ctime     int    `json:"ctime"`                // 1575483142
	Desc      string `json:"desc"`                 // "本家：sm32825363 ...
	State     int    `json:"state"`                // 0
	Attribute int    `json:"attribute"`            // 16793984
	Duration  int    `json:"duration"`             // 268
	MissionID int    `json:"mission_id,omitempty"` // 271861
	Rights    struct {
		Bp            int `json:"bp"`              // 0
		Elec          int `json:"elec"`            // 0
		Download      int `json:"download"`        // 0
		Movie         int `json:"movie"`           // 0
		Pay           int `json:"pay"`             // 0
		Hd5           int `json:"hd5"`             // 1
		NoReprint     int `json:"no_reprint"`      // 1
		Autoplay      int `json:"autoplay"`        // 1
		UgcPay        int `json:"ugc_pay"`         // 0
		IsCooperation int `json:"is_cooperation"`  // 1
		UgcPayPreview int `json:"ugc_pay_preview"` // 0
		NoBackground  int `json:"no_background"`   // 0
		ArcPay        int `json:"arc_pay"`         // 0
		PayFreeWatch  int `json:"pay_free_watch"`  // 0
	} `json:"rights"`
	Owner struct {
		MID  int    `json:"mid"`  // 434334701
		Name string `json:"name"` // "七海Nana7mi"
		Face string `json:"face"` // "https://i2.hdslb.com/bfs/face/3adb26401cfab0fe6b1a0d5b2c09220499108d64.jpg"
	} `json:"owner"`
	Stat struct {
		AID      int `json:"aid"`      // 78090377
		View     int `json:"view"`     // 403958
		Danmaku  int `json:"danmaku"`  // 1325
		Reply    int `json:"reply"`    // 1181
		Favorite int `json:"favorite"` // 10626
		Coin     int `json:"coin"`     // 13492
		Share    int `json:"share"`    // 1273
		NowRank  int `json:"now_rank"` // 0
		HisRank  int `json:"his_rank"` // 0
		Like     int `json:"like"`     // 18030
		Dislike  int `json:"dislike"`  // 0
		Vt       int `json:"vt"`       // 518998
		Vv       int `json:"vv"`       // 403958
		FavG     int `json:"fav_g"`    // 0
		LikeG    int `json:"like_g"`   // 0
	} `json:"stat"`
	Dynamic   string `json:"dynamic"` // "脆脆鲨组从我关注数到1w的时候就开始企划的手书 ...
	CID       int    `json:"cid"`     // 133606284
	Dimension struct {
		Width  int `json:"width"`  // 1920
		Height int `json:"height"` // 1080
		Rotate int `json:"rotate"` // 0
	} `json:"dimension"`
	ShortLinkV2      string `json:"short_link_v2"`          // "https://b23.tv/BV1vJ411B7ng"
	UpFromV2         int    `json:"up_from_v2,omitempty"`   // 35
	FirstFrame       string `json:"first_frame"`            // "http://i1.hdslb.com/bfs/storyff/n211202qn8mrsxajbj4fpwsrexvbo7ju_firsti.jpg"
	PubLocation      string `json:"pub_location,omitempty"` // "广东"
	VtDisplay        string `json:"vt_display"`             // "51.9万"
	Cover43          string `json:"cover43"`                // ""
	TIDv2            int    `json:"tidv2"`                  // 2020
	Tnamev2          string `json:"tnamev2"`                // "翻唱"
	PidV2            int    `json:"pid_v2"`                 // 1008
	PidNameV2        string `json:"pid_name_v2"`            // "游戏"
	BVID             string `json:"bvid"`                   // "BV1vJ411B7ng"
	Coins            int    `json:"coins"`                  // 2
	Time             int    `json:"time"`                   // 1745738818
	IP               string `json:"ip"`                     // ""
	InterVideo       bool   `json:"inter_video"`            // false
	ResourceType     string `json:"resource_type"`          // "ugc"
	Subtitle         string `json:"subtitle"`               // ""
	EnableVt         int    `json:"enable_vt"`              // 0
	SeasonID         int    `json:"season_id,omitempty"`    // 2832677
	Reason           string `json:"reason"`                 // ""
	IsChargingArc    bool   `json:"is_charging_arc"`        // false
	ElecArcType      int    `json:"elec_arc_type"`          // 0
	PlaybackPosition int    `json:"playback_position"`      // 0
	IsSelfView       bool   `json:"is_self_view"`           // false
}

// 查询用户置顶视频
type Arc struct {
	req.Get

	// 目标用户 UID
	Vmid int `api:"query"`
}

func (Arc) RawURL() string {
	return "/space/top/arc"
}

type ArcResponse struct {
	Error
	Data VideoInfo `json:"data"`
}

// 查询用户置顶视频
func GetArc(uid int) (result ArcResponse, err error) {
	err = cli.Result(Arc{Vmid: uid}, &result)
	return
}

// 设置置顶视频
type Set struct {
	PostCSRF
	*Credential

	// 置顶目标稿件 avid
	// 与 bvid 任选一个
	AID int `api:"body"`

	// 置顶目标稿件 bvid
	// 与 avid 任选一个
	BVID string `api:"body" req:"bvid"`

	// 置顶视频备注 最大 40 字符
	Reason string `api:"body"`
}

func (Set) RawURL() string {
	return "/space/top/arc/set"
}

type SetResponse struct {
	Error
}

// 设置置顶视频
func PostSet(bvid, reason string, credential *Credential) (result SetResponse, err error) {
	err = cli.Result(Set{BVID: bvid, Reason: reason, Credential: credential}, &result)
	return
}

// 取消置顶视频
type Cancel struct {
	PostCSRF
	*Credential
}

func (Cancel) RawURL() string {
	return "/space/top/arc/cancel"
}

type CancelResponse struct {
	Error
}

// 取消置顶视频
func PostCancel(credential *Credential) (result CancelResponse, err error) {
	err = cli.Result(Cancel{Credential: credential}, &result)
	return
}

// 查询用户代表作视频列表
type Masterpiece struct {
	req.Get

	// 目标用户 UID
	Vmid int `api:"query"`
}

func (Masterpiece) RawURL() string {
	return "/space/masterpiece"
}

type MasterpieceResponse struct {
	Error
	Data []VideoInfo `json:"data"`
}

// 查询用户代表作视频列表
func GetMasterpiece(uid int) (result MasterpieceResponse, err error) {
	err = cli.Result(Masterpiece{Vmid: uid}, &result)
	return
}

// 添加代表作视频
type Add struct {
	PostCSRF
	*Credential

	// 置顶目标稿件 avid
	// 与 bvid 任选一个
	AID int `api:"body"`

	// 置顶目标稿件 bvid
	// 与 avid 任选一个
	BVID string `api:"body" req:"bvid"`

	// 置顶视频备注 最大 40 字符
	Reason string `api:"body"`
}

func (Add) RawURL() string {
	return "/space/masterpiece/add"
}

type AddResponse struct {
	Error
}

// 添加代表作视频
func PostAdd(bvid, reason string, credential *Credential) (result AddResponse, err error) {
	err = cli.Result(Add{BVID: bvid, Reason: reason, Credential: credential}, &result)
	return
}

// 删除代表作视频
type MasterpieceCancel struct {
	PostCSRF
	*Credential

	// 取消置顶稿件 avid
	// 与 bvid 任选一个
	AID int `api:"body"`

	// 取消置顶稿件 bvid
	// 与 avid 任选一个
	BVID string `api:"body" req:"bvid"`
}

func (MasterpieceCancel) RawURL() string {
	return "/space/masterpiece/cancel"
}

type MasterpieceCancelResponse struct {
	Error
}

// 删除代表作视频
func PostMasterpieceCancel(bvid string, credential *Credential) (result MasterpieceCancelResponse, err error) {
	err = cli.Result(MasterpieceCancel{BVID: bvid, Credential: credential}, &result)
	return
}

// 查看用户个人 TAG
type AccTags struct {
	req.Get

	// 用户 UID
	MID int `api:"query"`
}

func (AccTags) RawURL() string {
	return "/space/acc/tags"
}

type AccTagsResponse struct {
	Error
	Data []struct {
		MID  int      `json:"mid"` // 434334701
		Tags []string `json:"tags"`
	} `json:"data"`
}

// 查看用户个人 TAG
func GetAccTags(uid int) (result AccTagsResponse, err error) {
	err = cli.Result(AccTags{MID: uid}, &result)
	return
}

// 修改个人 TAG
type TagsSet struct {
	PostCSRF
	*Credential

	// 要设置的 TAG 内容 用(,)间隔
	Tags string `api:"body"`
}

func (TagsSet) RawURL() string {
	return "/space/acc/tags/set"
}

type TagsSetResponse struct {
	Error
}

// 修改个人 TAG
func PostTagsSet(tags []string, credential *Credential) (result TagsSetResponse, err error) {
	err = cli.Result(TagsSet{Tags: strings.Join(tags, ","), Credential: credential}, &result)
	return
}

// 查看用户空间公告
type Notice struct {
	req.Get

	// 用户 UID
	MID int `api:"query"`
}

func (Notice) RawURL() string {
	return "/space/notice"
}

type NoticeResponse struct {
	Error
	Data string `json:"data"` // "新浪微博：七海Nana7mi"
}

// 查看用户空间公告
func GetNotice(uid int) (result NoticeResponse, err error) {
	err = cli.Result(Notice{MID: uid}, &result)
	return
}

// 修改空间公告
type NoticeSet struct {
	PostCSRF
	*Credential

	// 要设置的公告内容
	Notice string `api:"body"`
}

func (NoticeSet) RawURL() string {
	return "/space/notice/set"
}

type NoticeSetResponse struct {
	Error
}

// 修改空间公告
func PostNoticeSet(notice string, credential *Credential) (result NoticeSetResponse, err error) {
	err = cli.Result(NoticeSet{Notice: notice, Credential: credential}, &result)
	return
}

// 查询空间设置
type Settings struct {
	req.Get

	// 用户 UID
	MID int `api:"query"`
}

func (Settings) RawURL() string {
	return "http://space.bilibili.com/ajax/settings/getSettings"
}

type SettingsResponse struct {
	Status bool `json:"status"` // true
	Data   struct {
		Privacy struct {
			Bangumi           int `json:"bangumi"`             // 0
			Bbq               int `json:"bbq"`                 // 1
			Channel           int `json:"channel"`             // 1
			ChargeVideo       int `json:"charge_video"`        // 1
			CloseSpaceMedal   int `json:"close_space_medal"`   // 0
			CoinsVideo        int `json:"coins_video"`         // 1
			Comic             int `json:"comic"`               // 0
			DisableFollowing  int `json:"disable_following"`   // 1
			DisableShowFans   int `json:"disable_show_fans"`   // 0
			DisableShowNft    int `json:"disable_show_nft"`    // 0
			DisableShowSchool int `json:"disable_show_school"` // 0
			DressUp           int `json:"dress_up"`            // 1
			FavVideo          int `json:"fav_video"`           // 0
			Groups            int `json:"groups"`              // 0
			LessonVideo       int `json:"lesson_video"`        // 1
			LikesVideo        int `json:"likes_video"`         // 1
			LivePlayback      int `json:"live_playback"`       // 0
			OnlyShowWearing   int `json:"only_show_wearing"`   // 1
			PlayedGame        int `json:"played_game"`         // 0
			Tags              int `json:"tags"`                // 0
			UserInfo          int `json:"user_info"`           // 1
		} `json:"privacy"`
		ShowNftSwitch bool `json:"show_nft_switch"` // true
		IndexOrder    []struct {
			ID   int    `json:"id"`   // 1
			Name string `json:"name"` // "我的稿件"
		} `json:"index_order"`
		Theme               string `json:"theme"`                  // "default"
		ThemePreviewImgPath string `json:"theme_preview_img_path"` // ""
		Toutu               struct {
			SID          int    `json:"sid"`           // 1
			Expire       int64  `json:"expire"`        // 2996777597
			SImg         string `json:"s_img"`         // "bfs/space/768cc4fd97618cf589d23c2711a1d1a729f42235.png"
			LImg         string `json:"l_img"`         // "bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png"
			AndroidImg   string `json:"android_img"`   // ""
			IphoneImg    string `json:"iphone_img"`    // ""
			IpadImg      string `json:"ipad_img"`      // ""
			ThumbnailImg string `json:"thumbnail_img"` // ""
			Platform     int    `json:"platform"`      // 0
		} `json:"toutu"`
	} `json:"data"`
}

// 查询空间设置
func GetSettings(uid int) (result SettingsResponse, err error) {
	err = cli.Result(Settings{MID: uid}, &result)
	return
}

// 查询可用头图列表 (Web端)
type TopPhotoList struct {
	req.Get

	// 用户 UID
	MID int `api:"query"`
}

func (TopPhotoList) RawURL() string {
	return "https://space.bilibili.com/ajax/topphoto/getlist"
}

type TopPhotoListResponse struct {
	Status bool `json:"status"` // true
	Data   []struct {
		ID           int    `json:"id"`            // 1
		ProductName  string `json:"product_name"`  // "bilibili春"
		Price        int    `json:"price"`         // 0
		CoinType     int    `json:"coin_type"`     // 0
		VipFree      int    `json:"vip_free"`      // 0
		SImg         string `json:"s_img"`         // "bfs/space/768cc4fd97618cf589d23c2711a1d1a729f42235.png"
		LImg         string `json:"l_img"`         // "bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png"
		ThumbnailImg string `json:"thumbnail_img"` // ""
		SortNum      int    `json:"sort_num"`      // 19
		IsDisable    int    `json:"is_disable"`    // 0
		Expire       int64  `json:"expire"`        // 2890206564
		Had          int    `json:"had"`           // 1
	} `json:"data"`
}

// 查询可用头图列表 (Web端)
func GetTopPhotoList(uid int) (result TopPhotoListResponse, err error) {
	err = cli.Result(TopPhotoList{MID: uid}, &result)
	return
}

// 设置空间头图 (Web端)
type SetToutu struct {
	PostCSRF
	*Credential

	// 头图 ID
	ID int `api:"body"`
}

func (SetToutu) RawURL() string {
	return "https://space.bilibili.com/ajax/settings/setToutu"
}

type SetToutuResponse struct {
	Status bool   `json:"status"`
	Data   string `json:"data"`
}

// 设置空间头图 (Web端)
func PostSetToutu(id int, credential *Credential) (result SetToutuResponse, err error) {
	err = cli.Result(SetToutu{ID: id, Credential: credential}, &result)
	return
}

// 查询用户最近玩过的游戏
type LastPlayGame struct {
	req.Get
	*Credential

	// 目标用户 UID
	MID int `api:"query"`
}

func (LastPlayGame) RawURL() string {
	return "/space/lastplaygame/v2"
}

type LastPlayGameResponse struct {
	Error
	Data struct {
		PageNum    int `json:"page_num"`    // 0
		PageSize   int `json:"page_size"`   // 15
		TotalCount int `json:"total_count"` // 1
		List       []struct {
			GameBaseID   int      `json:"game_base_id"` // 102216
			GameName     string   `json:"game_name"`    // "公主连结Re:Dive"
			GameIcon     string   `json:"game_icon"`    // "https://i0.hdslb.com/bfs/game/3bb819e010fe6d594d8f4d417ee380f40e8b5b06.png"
			Grade        float64  `json:"grade"`        // 8.5
			DetailURL    string   `json:"detail_url"`   // "https://www.biligame.com/detail/?id=102216"
			GameTags     []string `json:"game_tags"`
			Notice       string   `json:"notice"`         // "五周年庆典开启中！登录就送10000宝石！"
			GiftTitle    string   `json:"gift_title"`     // ""
			GameStatusV2 int      `json:"game_status_v2"` // 0
		} `json:"list"`
	} `json:"data"`
}

// 查询用户最近玩过的游戏
func GetLastPlayGame(uid int, credential *Credential) (result LastPlayGameResponse, err error) {
	err = cli.Result(LastPlayGame{MID: uid, Credential: credential}, &result)
	return
}

// 查询用户最近投币视频（Web）
type CoinVideo struct {
	req.Get
	*Credential

	// 目标用户 UID
	Vmid int `api:"query"`
}

func (CoinVideo) RawURL() string {
	return "/space/coin/video"
}

type CoinVideoResponse struct {
	Error
	Data []VideoInfo `json:"data"`
}

// 查询用户最近投币视频（Web）
func GetCoinVideo(uid int, credential *Credential) (result CoinVideoResponse, err error) {
	err = cli.Result(CoinVideo{Vmid: uid, Credential: credential}, &result)
	return
}

// 查询用户最近点赞视频（Web）
type LikeVideo struct {
	req.Get
	*Credential

	// 目标用户 UID
	Vmid int `api:"query"`
}

func (LikeVideo) RawURL() string {
	return "/space/like/video"
}

type LikeVideoResponse struct {
	Error
	Data struct {
		List []VideoInfo `json:"list"`
	} `json:"data"`
}

// 查询用户最近点赞视频（Web）
func GetLikeVideo(uid int, credential *Credential) (result LikeVideoResponse, err error) {
	err = cli.Result(LikeVideo{Vmid: uid, Credential: credential}, &result)
	return
}

// 查询用户投稿视频明细
type ArcSearch struct {
	GetWBI
	*Credential

	// 目标用户 UID
	MID int `api:"query"`

	// 排序方式
	// 默认为 pubdate
	// 最新发布 pubdate
	// 最多播放 click
	// 最多收藏 stow
	Order string `api:"query"`

	// 筛选目标分区
	// 默认为 0 不进行分区筛选
	TID int `api:"query"`

	// 关键词筛选
	// 用于使用关键词搜索该UP主视频稿件
	Keyword string `api:"query"`

	// 页码
	// 默认为 1
	Pn int `api:"query:1"`

	// 每页项数
	// 默认为 30
	Ps int `api:"query:30"`
}

func (ArcSearch) RawURL() string {
	return "/space/wbi/arc/search"
}

type ArcSearchResponse struct {
	Error
	Data struct {
		List struct {
			Slist []any `json:"slist"`
			Tlist map[int]struct {
				TID   int    `json:"tid"`   // 4
				Count int    `json:"count"` // 13
				Name  string `json:"name"`  // "游戏"
			} `json:"tlist"`
			Vlist []struct {
				Comment          int    `json:"comment"`            // 44
				TypeID           int    `json:"typeid"`             // 27
				Play             int    `json:"play"`               // 5258
				Pic              string `json:"pic"`                // "http://i1.hdslb.com/bfs/archive/090bdea9fa9e5cacd78f50961e4db615d13cee5e.jpg"
				Subtitle         string `json:"subtitle"`           // ""
				Description      string `json:"description"`        // "伊蕾娜生贺作品 不准骂我"
				Copyright        string `json:"copyright"`          // "1"
				Title            string `json:"title"`              // "我是灰之魔女伊蕾娜"
				Review           int    `json:"review"`             // 0
				Author           string `json:"author"`             // "脆鲨12138"
				MID              int    `json:"mid"`                // 188888131
				Created          int    `json:"created"`            // 1729094460
				Length           string `json:"length"`             // "00:18"
				VideoReview      int    `json:"video_review"`       // 2
				AID              int64  `json:"aid"`                // 113315128938366
				BVID             string `json:"bvid"`               // "BV1hxmwYDEJ6"
				HideClick        bool   `json:"hide_click"`         // false
				IsPay            int    `json:"is_pay"`             // 0
				IsUnionVideo     int    `json:"is_union_video"`     // 0
				IsSteinsGate     int    `json:"is_steins_gate"`     // 0
				IsLivePlayback   int    `json:"is_live_playback"`   // 0
				IsLessonVideo    int    `json:"is_lesson_video"`    // 0
				IsLessonFinished int    `json:"is_lesson_finished"` // 0
				LessonUpdateInfo string `json:"lesson_update_info"` // ""
				JumpURL          string `json:"jump_url"`           // ""
				Meta             any    `json:"meta"`
				IsAvoided        int    `json:"is_avoided"`        // 0
				SeasonID         int    `json:"season_id"`         // 0
				Attribute        int    `json:"attribute"`         // 16640
				IsChargingArc    bool   `json:"is_charging_arc"`   // false
				ElecArcType      int    `json:"elec_arc_type"`     // 0
				Vt               int    `json:"vt"`                // 0
				EnableVt         int    `json:"enable_vt"`         // 0
				VtDisplay        string `json:"vt_display"`        // ""
				PlaybackPosition int    `json:"playback_position"` // 0
				IsSelfView       bool   `json:"is_self_view"`      // false
			} `json:"vlist"`
		} `json:"list"`
		Page struct {
			Pn    int `json:"pn"`    // 1
			Ps    int `json:"ps"`    // 30
			Count int `json:"count"` // 77
		} `json:"page"`
		EpisodicButton struct {
			Text string `json:"text"` // "播放全部"
			URI  string `json:"uri"`  // "//www.bilibili.com/medialist/play/188888131?from=space"
		} `json:"episodic_button"`
		IsRisk      bool `json:"is_risk"`       // false
		GaiaResType int  `json:"gaia_res_type"` // 0
		GaiaData    any  `json:"gaia_data"`
	} `json:"data"`
}

// 查询用户投稿视频明细
func GetArcSearch(uid int, credential *Credential) (result ArcSearchResponse, err error) {
	err = cli.Result(ArcSearch{MID: uid, Credential: credential}, &result)
	return
}

// 查询用户投稿专栏明细
type Article struct {
	GetWBI
	*Credential

	// 目标用户 UID
	MID int `api:"query"`

	// 页码
	// 默认为 1
	Pn int `api:"query:1"`

	// 每页项数
	// 默认为 30
	Ps int `api:"query:30"`

	// 排序方式
	// 默认为 publish_time
	// 最新发布 publish_time
	// 最多阅读 view
	// 最多收藏 fav
	Sort string `api:"query"`
}

func (Article) RawURL() string {
	return "/space/wbi/article"
}

type ArticleResponse struct {
	Error
	Data struct {
		Articles []struct {
			ID       int `json:"id"` // 41510926
			Category struct {
				ID       int    `json:"id"`        // 15
				ParentID int    `json:"parent_id"` // 3
				Name     string `json:"name"`      // "日常"
			} `json:"category"`
			Categories []struct {
				ID       int    `json:"id"`        // 3
				ParentID int    `json:"parent_id"` // 0
				Name     string `json:"name"`      // "生活"
			} `json:"categories"`
			Title      string `json:"title"`       // "B站UP主日报2025年04月29日23点（v1.1）"
			Summary    string `json:"summary"`     // "粉丝数第3多：知识 罗翔说刑法 (3188.61万粉, +2541)昨日涨粉第1多：美食 灶台阿连 (+2.17万粉)昨日掉粉最多：游戏 阿Q君的吉拔猫 (-11376粉)昨日涨播放第3多：资讯 河南都市频道 (+1014.60万)昨日点赞数第1多：游戏 明日方舟 (+34.50万)昨日涨阅读第1多：运动 茉莉蜜糕 (+10.59万)昨日充电人次最多：美食 食贫道 (+308人)各位观众早上好，今天是2025年04月30日，农历四月初三，欢迎收看狸子每天定期更新的新人联播节目 o(∩_∩)o近期UP主宏观"
			BannerURL  string `json:"banner_url"`  // ""
			TemplateID int    `json:"template_id"` // 4
			State      int    `json:"state"`       // 0
			Author     struct {
				MID     int    `json:"mid"`  // 300021061
				Name    string `json:"name"` // "狸工智能"
				Face    string `json:"face"` // "https://i1.hdslb.com/bfs/face/4cba9bc9d6cf6935a37ec156dedb8f8d26c1df95.jpg"
				Pendant struct {
					PID    int    `json:"pid"`    // 0
					Name   string `json:"name"`   // ""
					Image  string `json:"image"`  // ""
					Expire int    `json:"expire"` // 0
				} `json:"pendant"`
				OfficialVerify struct {
					Type int    `json:"type"` // -1
					Desc string `json:"desc"` // ""
				} `json:"official_verify"`
				Nameplate struct {
					NID        int    `json:"nid"`         // 3
					Name       string `json:"name"`        // "白银殿堂"
					Image      string `json:"image"`       // "https://i2.hdslb.com/bfs/face/f6a31275029365ae5dc710006585ddcf1139bde1.png"
					ImageSmall string `json:"image_small"` // "https://i1.hdslb.com/bfs/face/b09cdb4c119c467cf2d15db5263b4f539fa6e30b.png"
					Level      string `json:"level"`       // "高级勋章"
					Condition  string `json:"condition"`   // "单个自制视频总播放数>=10万"
				} `json:"nameplate"`
				Vip struct {
					Type       int `json:"type"`         // 1
					Status     int `json:"status"`       // 0
					DueDate    int `json:"due_date"`     // 0
					VipPayType int `json:"vip_pay_type"` // 0
					ThemeType  int `json:"theme_type"`   // 0
					Label      struct {
						Path       string `json:"path"`        // ""
						Text       string `json:"text"`        // ""
						LabelTheme string `json:"label_theme"` // ""
					} `json:"label"`
					AvatarSubscript int    `json:"avatar_subscript"` // 0
					NicknameColor   string `json:"nickname_color"`   // ""
				} `json:"vip"`
			} `json:"author"`
			Reprint     int      `json:"reprint"` // 0
			ImageUrls   []string `json:"image_urls"`
			PublishTime int      `json:"publish_time"` // 1745981640
			Ctime       int      `json:"ctime"`        // 1745981640
			Mtime       int      `json:"mtime"`        // 1745981640
			Stats       struct {
				View     int `json:"view"`     // 1927
				Favorite int `json:"favorite"` // 0
				Like     int `json:"like"`     // 24
				Dislike  int `json:"dislike"`  // 0
				Reply    int `json:"reply"`    // 0
				Share    int `json:"share"`    // 0
				Coin     int `json:"coin"`     // 0
				Dynamic  int `json:"dynamic"`  // 0
			} `json:"stats"`
			Words           int      `json:"words"` // 0
			OriginImageUrls []string `json:"origin_image_urls"`
			List            any      `json:"list"`
			IsLike          bool     `json:"is_like"` // false
			Media           struct {
				Score    int    `json:"score"`     // 0
				MediaID  int    `json:"media_id"`  // 0
				Title    string `json:"title"`     // ""
				Cover    string `json:"cover"`     // ""
				Area     string `json:"area"`      // ""
				TypeID   int    `json:"type_id"`   // 0
				TypeName string `json:"type_name"` // ""
				Spoiler  int    `json:"spoiler"`   // 0
			} `json:"media"`
			ApplyTime        string `json:"apply_time"` // ""
			CheckTime        string `json:"check_time"` // ""
			Original         int    `json:"original"`   // 1
			ActID            int    `json:"act_id"`     // 0
			Dispute          any    `json:"dispute"`
			AuthenMark       any    `json:"authenMark"`
			CoverAvid        int    `json:"cover_avid"` // 0
			TopVideoInfo     any    `json:"top_video_info"`
			Type             int    `json:"type"`                 // 0
			CheckState       int    `json:"check_state"`          // 0
			OriginTemplateID int    `json:"origin_template_id"`   // 5
			Attributes       int    `json:"attributes,omitempty"` // 148
		} `json:"articles"`
		Pn    int `json:"pn"`    // 1
		Ps    int `json:"ps"`    // 30
		Count int `json:"count"` // 2489
	} `json:"data"`
}

// 查询用户投稿专栏明细
func GetArticle(uid int, credential *Credential) (result ArticleResponse, err error) {
	err = cli.Result(Article{MID: uid, Credential: credential}, &result)
	return
}

// 查询用户专栏文集明细
type ArticleLists struct {
	req.Get
	*Credential

	// 目标用户 UID
	MID int `api:"query"`

	// 排序方式
	// 默认为 0
	// 最近更新 0
	// 最多阅读 1
	Sort int `api:"query"`
}

func (ArticleLists) RawURL() string {
	return "/article/up/lists"
}

type ArticleListsResponse struct {
	Error
	Data struct {
		Lists []struct {
			ID            int    `json:"id"`             // 26407
			MID           int    `json:"mid"`            // 2859372
			Name          string `json:"name"`           // "周榜"
			ImageURL      string `json:"image_url"`      // "https://i0.hdslb.com/bfs/article/96d2b3d2a72e6497a011c885ab9245c51507ce18.png"
			UpdateTime    int    `json:"update_time"`    // 1745768376
			Ctime         int    `json:"ctime"`          // 1537942450
			PublishTime   int    `json:"publish_time"`   // 1745768458
			Summary       string `json:"summary"`        // ""
			Words         int    `json:"words"`          // 162765
			Read          int    `json:"read"`           // 55220355
			ArticlesCount int    `json:"articles_count"` // 339
			State         int    `json:"state"`          // 1
			Reason        string `json:"reason"`         // ""
			ApplyTime     string `json:"apply_time"`     // ""
			CheckTime     string `json:"check_time"`     // ""
		} `json:"lists"`
		Total int `json:"total"` // 10
	} `json:"data"`
}

// 查询用户专栏文集明细
func GetArticleLists(uid int, credential *Credential) (result ArticleListsResponse, err error) {
	err = cli.Result(ArticleLists{MID: uid, Credential: credential}, &result)
	return
}

// 查询用户投稿音频明细
type SongUpper struct {
	req.Get

	// 目标用户 UID
	UID int `api:"query"`

	// 页码
	// 默认为 1
	Pn int `api:"query:1"`

	// 每页项数
	// 默认为 30
	Ps int `api:"query:30"`

	// 排序方式
	// 最新发布 1
	// 最多播放 2
	// 最多收藏 3
	Order string `api:"query"`
}

func (SongUpper) RawURL() string {
	return "https://api.bilibili.com/audio/music-service/web/song/upper"
}

type SongUpperResponse struct {
	Error
	Msg  string `json:"msg"` // "success"
	Data struct {
		CurPage   int `json:"curPage"`   // 1
		PageCount int `json:"pageCount"` // 1
		TotalSize int `json:"totalSize"` // 2
		PageSize  int `json:"pageSize"`  // 2
		Data      []struct {
			ID         int    `json:"id"`         // 378521
			UID        int    `json:"uid"`        // 8047632
			Uname      string `json:"uname"`      // "哔哩哔哩弹幕网"
			Author     string `json:"author"`     // ""
			Title      string `json:"title"`      // "《B TOGETHER》-bilibili九周年主题曲"
			Cover      string `json:"cover"`      // "http://i0.hdslb.com/bfs/music/109136c63e16d83fbad5ec9282a6fb96498d8144.jpg"
			Intro      string `json:"intro"`      // ""
			Lyric      string `json:"lyric"`      // "http://i0.hdslb.com/bfs/music/1529979007378521.lrc"
			Crtype     int    `json:"crtype"`     // 1
			Duration   int    `json:"duration"`   // 0
			Passtime   int    `json:"passtime"`   // 1529928347
			Curtime    int    `json:"curtime"`    // 0
			AID        int    `json:"aid"`        // 0
			BVID       string `json:"bvid"`       // ""
			CID        int    `json:"cid"`        // 0
			MsID       int    `json:"msid"`       // 0
			Attr       int    `json:"attr"`       // 0
			Limit      int    `json:"limit"`      // 0
			ActivityID int    `json:"activityId"` // 0
			Limitdesc  string `json:"limitdesc"`  // ""
			CoinNum    int    `json:"coin_num"`   // 3640
			Ctime      int64  `json:"ctime"`      // 1529928235000
			Statistic  struct {
				SID     int `json:"sid"`     // 378521
				Play    int `json:"play"`    // 123582
				Collect int `json:"collect"` // 5510
				Comment int `json:"comment"` // 1591
				Share   int `json:"share"`   // 535
			} `json:"statistic"`
			VipInfo    any `json:"vipInfo"`
			CollectIds any `json:"collectIds"`
			IsCooper   int `json:"is_cooper"` // 0
		} `json:"data"`
	} `json:"data"`
}

// 查询用户投稿音频明细
func GetSongUpper(uid int) (result SongUpperResponse, err error) {
	err = cli.Result(SongUpper{UID: uid}, &result)
	return
}

// 做吐了 先做到这里吧
// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/space.html#%E9%A2%91%E9%81%93
