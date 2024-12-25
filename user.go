package api

import (
	"strings"

	"github.com/Drelf2018/req"
)

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/info.html

// 用户空间详细信息
type AccInfo struct {
	GetWBI
	AccessID
	*Credential

	// 目标用户 UID
	MID int `api:"query"`
}

func (AccInfo) RawURL() string {
	return "/space/wbi/acc/info"
}

type AccInfoResponse struct {
	Error
	TTL  int `json:"ttl"` // 1
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
			Pid               int    `json:"pid"`                 // 0
			Name              string `json:"name"`                // ""
			Image             string `json:"image"`               // ""
			Expire            int    `json:"expire"`              // 0
			ImageEnhance      string `json:"image_enhance"`       // ""
			ImageEnhanceFrame string `json:"image_enhance_frame"` // ""
			NPid              int    `json:"n_pid"`               // 0
		} `json:"pendant"`
		Nameplate struct {
			Nid        int    `json:"nid"`         // 0
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
			Roomid        int    `json:"roomid"`         // 3394945
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
	TTL  int `json:"ttl"` // 1
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
				Pid               int    `json:"pid"`                 // 0
				Name              string `json:"name"`                // ""
				Image             string `json:"image"`               // ""
				Expire            int    `json:"expire"`              // 0
				ImageEnhance      string `json:"image_enhance"`       // ""
				ImageEnhanceFrame string `json:"image_enhance_frame"` // ""
				NPid              int    `json:"n_pid"`               // 0
			} `json:"pendant"`
			Nameplate struct {
				Nid        int    `json:"nid"`         // 0
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
	err = cli.Result(Card{MID: uid, Credential: credential, Photo: true}, &result)
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
	TTL  int `json:"ttl"` // 1
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
			Pid               int    `json:"pid"`                 // 4820
			Name              string `json:"name"`                // "七海演唱会"
			Image             string `json:"image"`               // "https://i1.hdslb.com/bfs/garb/item/7061b098537482bb78242a16cba5d44b426b429d.png"
			Expire            int    `json:"expire"`              // 0
			ImageEnhance      string `json:"image_enhance"`       // "https://i1.hdslb.com/bfs/garb/item/7061b098537482bb78242a16cba5d44b426b429d.png"
			ImageEnhanceFrame string `json:"image_enhance_frame"` // ""
			NPid              int    `json:"n_pid"`               // 4820
		} `json:"pendant"`
		Nameplate struct {
			Nid        int    `json:"nid"`         // 62
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

	// 目标用户的 UID 列表
	// 每个成员间用 , 分隔
	// 最多200个成员
	UIDs string `api:"query"`
}

func (CardMap) RawURL() string {
	return "/polymer/pc-electron/v1/user/cards"
}

type CardMapResponse struct {
	Error
	TTL  int `json:"ttl"` // 1
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
func GetCardMap(credential *Credential, uid ...string) (result CardMapResponse, err error) {
	err = cli.Result(CardMap{Credential: credential, UIDs: strings.Join(uid, ",")}, &result)
	return
}

// 多用户详细信息切片
type CardSlice struct {
	req.Get
	*Credential

	// 目标用户的 UID 列表
	// 每个成员间用 , 分隔
	// 最多 50 个成员
	UIDs string `api:"query"`
}

func (CardSlice) RawURL() string {
	return "https://api.vc.bilibili.com/account/v1/user/cards"
}

type CardSliceResponse struct {
	Error
	TTL  int `json:"ttl"` // 1
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
func GetCardSlice(credential *Credential, uid ...string) (result CardSliceResponse, err error) {
	err = cli.Result(CardSlice{Credential: credential, UIDs: strings.Join(uid, ",")}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/status_number.html

// 关系状态数
type RelationStat struct {
	req.Get
	*Credential

	// APP 登录 Token
	AccessKey string `api:"query,omitempty"`

	// 目标用户 UID
	Vmid int `api:"query"`
}

func (RelationStat) RawURL() string {
	return "/relation/stat"
}

type RelationStatResponse struct {
	Error
	TTL  int `json:"ttl"` // 1
	Data struct {
		Mid       int `json:"mid"`       // 434334701
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

	// APP 登录 Token
	AccessKey string `api:"query,omitempty"`

	// 目标用户 UID
	MID int `api:"query"`
}

func (UPStat) RawURL() string {
	return "/space/upstat"
}

type UPStatResponse struct {
	Error
	TTL  int `json:"ttl"` // 1
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
	TTL  int `json:"ttl"` // 1
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

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/space.html

// https://socialsisteryi.github.io/bilibili-API-collect/docs/user/medals.html

type MedalWall struct {
	req.Get
	*Credential

	//目标 mid
	TargetID int `api:"query"`
}

func (MedalWall) RawURL() string {
	return "https://api.live.bilibili.com/xlive/web-ucenter/user/MedalWall"
}

type MedalWallResponse struct {
	Error
	TTL  int `json:"ttl"` // 1
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
				Ruid               int    `json:"ruid"`                  // 7706705
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

func GetMedalWall(uid int, credential *Credential) (result MedalWallResponse, err error) {
	err = cli.Result(MedalWall{TargetID: uid, Credential: credential}, &result)
	return
}
