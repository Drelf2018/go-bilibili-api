package api

import "github.com/Drelf2018/req"

// https://socialsisteryi.github.io/bilibili-API-collect/docs/vip/info.html

// 卡券状态查询
type My struct {
	req.Get
	*Credential
}

func (My) RawURL() string {
	return "/vip/privilege/my"
}

type MyResponse struct {
	Error
	Data struct {
		List []struct {
			Type            int    `json:"type"`              // 1
			State           int    `json:"state"`             // 0
			ExpireTime      int    `json:"expire_time"`       // 0
			VipType         int    `json:"vip_type"`          // 2
			NextReceiveDays int    `json:"next_receive_days"` // 21
			PeriodEndUnix   int    `json:"period_end_unix"`   // 1737129600
			IsCount         bool   `json:"is_count"`          // true
			Name            string `json:"name"`              // ""
			CouponCode      string `json:"coupon_code"`       // ""
			AppDescribe     string `json:"app_describe"`      // ""
			ReciveState     int    `json:"recive_state"`      // 1
			SalaryType      int    `json:"salary_type"`       // 0
			ExpParams       struct {
				ExpGroupTag string `json:"exp_group_tag"` // "45476"
				HitValue    int    `json:"hit_value"`     // 2
			} `json:"exp_params"`
			ExtraParams any `json:"extra_params"`
		} `json:"list"`
		IsShortVip     bool   `json:"is_short_vip"`     // false
		IsFreightOpen  bool   `json:"is_freight_open"`  // true
		Level          int    `json:"level"`            // 6
		CurExp         int    `json:"cur_exp"`          // 48579
		NextExp        int    `json:"next_exp"`         // -1
		IsVip          bool   `json:"is_vip"`           // true
		IsSeniorMember int    `json:"is_senior_member"` // 0
		Format060102   int    `json:"format060102"`     // 241218
		IsOverdueVip   bool   `json:"is_overdue_vip"`   // false
		VipStatus      int    `json:"vip_status"`       // 1
		VipType        int    `json:"vip_type"`         // 2
		KeeptimeEnd    int    `json:"keeptime_end"`     // 1737129600
		VipDueDate     int    `json:"vip_due_date"`     // 1797264000
		VipIsAnnual    bool   `json:"vip_is_annual"`    // true
		VipIsMonth     bool   `json:"vip_is_month"`     // false
		VipIsNewUser   bool   `json:"vip_is_new_user"`  // false
		BindPhone      string `json:"bind_phone"`
		TaobaoAccount  any    `json:"taobao_account"`
	} `json:"data"`
}

// 卡券状态查询
func GetMy(credential *Credential) (result MyResponse, err error) {
	err = cli.Result(My{Credential: credential}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/vip/center.html

// 大会员中心信息
type VIPCenterCombine struct {
	req.Get
	*Credential

	// 平台 web端(web) 安卓APP(android)
	Platform string `api:"query,omitempty"`

	Build   int    `api:"query,omitempty"`
	MobiApp string `api:"query,omitempty"`
}

func (VIPCenterCombine) RawURL() string {
	return "/vip/web/vip_center/combine"
}

type VIPCenterCombineResponse struct {
	Error
	Data struct {
		User struct {
			Account struct {
				MID            int    `json:"mid"`              // 188888131
				Name           string `json:"name"`             // "脆鲨12138"
				Sex            string `json:"sex"`              // "女"
				Face           string `json:"face"`             // "https://i1.hdslb.com/bfs/face/86faab4844dd2c45870fdafa8f2c9ce7be3e999f.jpg"
				Sign           string `json:"sign"`             // "虚拟艺人团体Vcslive成员，喜欢宅在家打游戏的脆鲨12138！！！"
				Rank           int    `json:"rank"`             // 10000
				Birthday       int    `json:"birthday"`         //
				IsFakeAccount  int    `json:"is_fake_account"`  // 0
				IsDeleted      int    `json:"is_deleted"`       // 0
				InRegAudit     int    `json:"in_reg_audit"`     // 0
				IsSeniorMember int    `json:"is_senior_member"` // 0
				NameRender     any    `json:"name_render"`
			} `json:"account"`
			Vip struct {
				MID        int   `json:"mid"`          // 188888131
				VipType    int   `json:"vip_type"`     // 2
				VipStatus  int   `json:"vip_status"`   // 1
				VipDueDate int64 `json:"vip_due_date"` // 1797264000000
				VipPayType int   `json:"vip_pay_type"` // 1
				ThemeType  int   `json:"theme_type"`   // 0
				Label      struct {
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
				AvatarSubscriptURL string `json:"avatar_subscript_url"` // "https://i0.hdslb.com/bfs/activity-plat/static/20231122/a5480f18ac08c1e30876e3fde84db784/r9dln6E2Uo.png"
				NicknameColor      string `json:"nickname_color"`       // "#FB7299"
				IsNewUser          bool   `json:"is_new_user"`          // false
				TipMaterial        any    `json:"tip_material"`
				VipIsAnnual        bool   `json:"vip_is_annual"`   // true
				VipIsMonth         bool   `json:"vip_is_month"`    // false
				IsAutoRenew        bool   `json:"is_auto_renew"`   // true
				VipIsValid         bool   `json:"vip_is_valid"`    // true
				VipIsOverdue       bool   `json:"vip_is_overdue"`  // false
				VipKeepTime        int    `json:"vip_keep_time"`   // 0
				VipExpireDays      int    `json:"vip_expire_days"` // 717
			} `json:"vip"`
			Tv struct {
				Type       int `json:"type"`         // 0
				VipPayType int `json:"vip_pay_type"` // 0
				Status     int `json:"status"`       // 0
				DueDate    int `json:"due_date"`     // 0
			} `json:"tv"`
			BackgroundImageSmall string `json:"background_image_small"` // ""
			BackgroundImageBig   string `json:"background_image_big"`   // ""
			PanelTitle           string `json:"panel_title"`            // "脆鲨12138"
			PanelSubTitle        string `json:"panel_sub_title"`        // ""
			AvatarPendant        struct {
				Image             string `json:"image"`               // "https://i0.hdslb.com/bfs/garb/item/7061b098537482bb78242a16cba5d44b426b429d.png"
				ImageEnhance      string `json:"image_enhance"`       // "https://i0.hdslb.com/bfs/garb/item/7061b098537482bb78242a16cba5d44b426b429d.png"
				ImageEnhanceFrame string `json:"image_enhance_frame"` // ""
			} `json:"avatar_pendant"`
			VipOverdueExplain    string `json:"vip_overdue_explain"`    // "年度大会员有效期 2026/12/15"
			TvOverdueExplain     string `json:"tv_overdue_explain"`     // "超级大会员未开通"
			AccountExceptionText string `json:"account_exception_text"` // ""
			IsAutoRenew          bool   `json:"is_auto_renew"`          // true
			IsTvAutoRenew        bool   `json:"is_tv_auto_renew"`       // false
			SurplusSeconds       int    `json:"surplus_seconds"`        // 61897858
			VipKeepTime          int    `json:"vip_keep_time"`          // 1600185600
			Renew                struct {
				Text string `json:"text"` // ""
				Link string `json:"link"` // ""
			} `json:"renew"`
			Notice struct {
				Text             string `json:"text"`               // ""
				TvText           string `json:"tv_text"`            // ""
				Type             int    `json:"type"`               // 0
				CanClose         bool   `json:"can_close"`          // false
				SurplusSeconds   int    `json:"surplus_seconds"`    // 0
				TvSurplusSeconds int    `json:"tv_surplus_seconds"` // 0
			} `json:"notice"`
			Face struct {
				ContainerSize struct {
					Width  float64 `json:"width"`  // 1.375
					Height float64 `json:"height"` // 1.375
				} `json:"container_size"`
				FallbackLayers struct {
					Layers []struct {
						Visible     bool `json:"visible"` // true
						GeneralSpec struct {
							PosSpec struct {
								CoordinatePos int     `json:"coordinate_pos"` // 2
								AxisX         float64 `json:"axis_x"`         // 0.6875
								AxisY         float64 `json:"axis_y"`         // 0.6875
							} `json:"pos_spec"`
							SizeSpec struct {
								Width  float64 `json:"width"`  // 0.787
								Height float64 `json:"height"` // 0.787
							} `json:"size_spec"`
							RenderSpec struct {
								Opacity int `json:"opacity"` // 1
							} `json:"render_spec"`
						} `json:"general_spec"`
						LayerConfig struct {
							Tags struct {
								AVATARLAYER struct {
								} `json:"AVATAR_LAYER"`
								GENERALCFG struct {
									ConfigType    int `json:"config_type"` // 1
									GeneralConfig struct {
										WebCSSStyle struct {
											BorderRadius string `json:"borderRadius"` // "50%"
										} `json:"web_css_style"`
									} `json:"general_config"`
								} `json:"GENERAL_CFG"`
							} `json:"tags"`
							IsCritical bool `json:"is_critical"` // true
						} `json:"layer_config"`
						Resource struct {
							ResType  int `json:"res_type"` // 3
							ResImage struct {
								ImageSrc struct {
									SrcType     int `json:"src_type"`    // 1
									Placeholder int `json:"placeholder"` // 6
									Remote      struct {
										URL      string `json:"url"`       // "https://i1.hdslb.com/bfs/face/86faab4844dd2c45870fdafa8f2c9ce7be3e999f.jpg"
										BfsStyle string `json:"bfs_style"` // "widget-layer-avatar"
									} `json:"remote"`
								} `json:"image_src"`
							} `json:"res_image"`
						} `json:"resource"`
					} `json:"layers"`
					IsCriticalGroup bool `json:"is_critical_group"` // true
				} `json:"fallback_layers"`
				MID string `json:"mid"` // "188888131"
			} `json:"face"`
		} `json:"user"`
		Wallet struct {
			Coupon            int  `json:"coupon"`             // 0
			Point             int  `json:"point"`              // 0
			PrivilegeReceived bool `json:"privilege_received"` // false
		} `json:"wallet"`
		UnionVip      any `json:"union_vip"`
		OtherOpenInfo any `json:"other_open_info"`
		Privileges    []struct {
			ID              int    `json:"id"`   // 3
			Name            string `json:"name"` // "视听特权"
			ChildPrivileges []struct {
				FirstID            int    `json:"first_id"`             // 3
				ReportID           string `json:"report_id"`            // "clearwatch"
				Name               string `json:"name"`                 // "超清看"
				Desc               string `json:"desc"`                 // "会员用户超清晰观看"
				Explain            string `json:"explain"`              // "大会员可专享高帧率、高码率画质（最高可达超清4k），觉醒超凡视觉体验。"
				IconURL            string `json:"icon_url"`             // "http://i0.hdslb.com/bfs/vip/b0faebe1b6b5178c49d6d9500c0ded88c0d677ad.png"
				IconGrayURL        string `json:"icon_gray_url"`        // ""
				BackgroundImageURL string `json:"background_image_url"` // "http://i0.hdslb.com/bfs/vip/21d79540e10618ee9bbaf8874ae711442d10edf0.png"
				Link               string `json:"link"`                 // "https://big.bilibili.com/mobile/explainDetails?name=%E8%B6%85%E6%B8%85%E7%9C%8B&closable=1&navhide=1"
				ImageURL           string `json:"image_url"`            // "http://i0.hdslb.com/bfs/vip/fbe4cf2288571d7b0a509c7014d5182789ffdd74.png"
				Type               int    `json:"type"`                 // 0
				HotType            int    `json:"hot_type"`             // 0
				NewType            int    `json:"new_type"`             // 1
				ID                 int    `json:"id"`                   // 3
			} `json:"child_privileges"`
		} `json:"privileges"`
		Banners any `json:"banners"`
		Welfare struct {
			Count int `json:"count"` // 6
			List  []struct {
				ID          int    `json:"id"`           // 120
				Name        string `json:"name"`         // "移动免费领4G流量"
				HomepageURI string `json:"homepage_uri"` // "https://i1.hdslb.com/bfs/vip/abec6a734d8ca1a2b8bf2cc036aafe28d43c0e99.png"
				BackdropURI string `json:"backdrop_uri"` // "https://i1.hdslb.com/bfs/vip/202df1b0786c66c2d5e41db1b048f0a8cdc219ac.png"
				TID         int    `json:"tid"`          // 0
				Rank        int    `json:"rank"`         // 1
				ReceiveURI  string `json:"receive_uri"`  // "https://wx.10086.cn/qwhdhub/leadin/1024120650?A_C_CODE=KXGxyMb60d"
			} `json:"list"`
		} `json:"welfare"`
		RecommendPendants struct {
			List []struct {
				ID      int    `json:"id"`       // 1758
				Name    string `json:"name"`     // "至尊戒"
				Image   string `json:"image"`    // "https://i0.hdslb.com/bfs/garb/item/025d07fa04d38236bc2258be2faf2867e2c48fe1.png"
				JumpURL string `json:"jump_url"` // "https://www.bilibili.com/h5/mall/preview/pendant/1758?navhide=1"
			} `json:"list"`
			JumpURL string `json:"jump_url"` // "https://www.bilibili.com/h5/mall/pendant/home?navhide=1&tab_id=22"
		} `json:"recommend_pendants"`
		RecommendCards struct {
			List []struct {
				ID      int    `json:"id"`       // 18
				Name    string `json:"name"`     // "阿维"
				Image   string `json:"image"`    // "https://i0.hdslb.com/bfs/vip/ffa3e8c1cf92eb0c01db61abe5741419e9302a70.png"
				JumpURL string `json:"jump_url"` // "https://www.bilibili.com/h5/mall/preview/feed/18?navhide=1"
			} `json:"list"`
			JumpURL string `json:"jump_url"` // "https://www.bilibili.com/h5/mall/card/detail?navhide=1&tab_id=5"
		} `json:"recommend_cards"`
		Sort []struct {
			Key  string `json:"key"`  // "union_vip"
			Sort int    `json:"sort"` // 1
		} `json:"sort"`
		InReview bool `json:"in_review"` // false
		BigPoint struct {
			PointInfo struct {
				Point       int `json:"point"`        // 0
				ExpirePoint int `json:"expire_point"` // 0
				ExpireTime  int `json:"expire_time"`  // 0
				ExpireDays  int `json:"expire_days"`  // 0
			} `json:"point_info"`
			SignInfo struct {
				SignRemind   bool `json:"sign_remind"`   // true
				Benefit      int  `json:"benefit"`       // 5
				BonusBenefit int  `json:"bonus_benefit"` // 200
				NormalRemind bool `json:"normal_remind"` // true
				MuggleTask   bool `json:"muggle_task"`   // true
				ExpValue     int  `json:"exp_value"`     // -1
			} `json:"sign_info"`
			SkuInfo struct {
				Skus []struct {
					Base struct {
						Token            string   `json:"token"`   // "1126774380639117732"
						Title            string   `json:"title"`   // "百草味50元礼盒直减券"
						Picture          string   `json:"picture"` // "https://i0.hdslb.com/bfs/activity-plat/9f7fe43bb6cd1692f10b9c0a2892ea63fc80041e.jpg"
						RotationPictures []string `json:"rotation_pictures"`
						Price            struct {
							Origin    int `json:"origin"` // 188
							Promotion any `json:"promotion"`
							Sale      int `json:"sale"` // 188
						} `json:"price"`
						Inventory struct {
							AvailableNum int `json:"available_num"` // 50000
							UsedNum      int `json:"used_num"`      // 44
							SurplusNum   int `json:"surplus_num"`   // 49956
						} `json:"inventory"`
						UserType          int `json:"user_type"`           // 2
						ExchangeLimitType int `json:"exchange_limit_type"` // 2
						ExchangeLimitNum  int `json:"exchange_limit_num"`  // 5
						StartTime         int `json:"start_time"`          // 1735198320
						EndTime           int `json:"end_time"`            // 1753977599
						State             int `json:"state"`               // 2
						Priority          int `json:"priority"`            // 90
					} `json:"base"`
				} `json:"skus"`
			} `json:"sku_info"`
			PointSwitchOff bool `json:"point_switch_off"` // false
			Tips           []struct {
				Content string `json:"content"` // "200大积分待领取"
			} `json:"tips"`
		} `json:"big_point"`
		HotList struct {
			Taps []struct {
				OID       string `json:"oid"`            // "{season_type:0,style_id:0,day:3,rank_type:5}"
				RankID    int    `json:"rank_id"`        // 279
				RankTitle string `json:"rank_title"`     // "会员"
				Type      int    `json:"type,omitempty"` // 1
			} `json:"taps"`
		} `json:"hot_list"`
		IntegrationTask bool `json:"integration_task"` // false
		FreeWelfare     []struct {
			ID             int    `json:"id"`             // 4290081
			Icon           string `json:"icon"`           // "https://i0.hdslb.com/bfs/bangumi/image/5a0027a21d5298fe62266968a0b65f1ab84efa72.png"
			Title          string `json:"title"`          // "开通联通王卡"
			TitleColor     string `json:"titleColor"`     // "#000000"
			Subtitle       string `json:"subtitle"`       // "每年享960G流量"
			SubtitleColor  string `json:"subtitleColor"`  // "#9499A0"
			Subtitle2      string `json:"subtitle2"`      // "领24个月大会员"
			Subtitle2Color string `json:"subtitle2Color"` // "#FF6699"
			TaskState      int    `json:"task_state"`     // 1
			Link           string `json:"link"`           // "https://www.bilibili.com/blackboard/activity-amrPj9oRBu.html?taskId=4290081&channel=cucc_king"
			Channel        string `json:"channel"`        // "cucc_king"
		} `json:"free_welfare"`
		ExtraParamas struct {
			SwitchOn            bool   `json:"switch_on"`               // false
			BirthdaySkuSwitchOn bool   `json:"birthday_sku_switch_on"`  // false
			IsBuyBirthdaySku    bool   `json:"is_buy_birthday_sku"`     // false
			IsBirthdayOff       bool   `json:"is_birthday_off"`         // false
			PhoneBindState      int    `json:"phone_bind_state"`        // 3
			AppTimes            int    `json:"app_times"`               // 0
			NaAb                int    `json:"na_ab"`                   // 0
			NaAbGroupID         string `json:"na_ab_group_id"`          // ""
			IsSameToLastSession bool   `json:"is_same_to_last_session"` // false
			IsWhiteList         bool   `json:"is_white_list"`           // false
			CloudAb             int    `json:"cloud_ab"`                // 0
			CloudAbGroupID      string `json:"cloud_ab_group_id"`       // ""
			BannerAb            int    `json:"banner_ab"`               // 0
			BannerAbGroupID     string `json:"banner_ab_group_id"`      // ""
			NowTime             int    `json:"now_time"`                // 0
			FreeWelfareAb       int    `json:"free_welfare_ab"`         // 0
			DeviceLimitAb       int    `json:"device_limit_ab"`         // 0
		} `json:"extra_paramas"`
		HitAb bool `json:"hit_ab"` // false
	} `json:"data"`
}

// 大会员中心信息
func GetVIPCenterCombine(credential *Credential) (result VIPCenterCombineResponse, err error) {
	err = cli.Result(VIPCenterCombine{Credential: credential}, &result)
	return
}

// 大积分中心信息
type HomepageCombine struct {
	req.Get
	*Credential
}

func (HomepageCombine) RawURL() string {
	return "https://api.biliapi.com/x/vip_point/homepage/combine"
}

type HomepageCombineResponse struct {
	Error
	Data struct {
		VipInfo struct {
			Type       int   `json:"type"`         // 2
			Status     int   `json:"status"`       // 1
			DueDate    int64 `json:"due_date"`     // 1797264000000
			VipPayType int   `json:"vip_pay_type"` // 1
			Label      struct {
				Path                  string `json:"path"`                      // "http://i0.hdslb.com/bfs/vip/label_annual.png"
				TextColor             string `json:"text_color"`                // ""
				BgStyle               int    `json:"bg_style"`                  // 0
				BgColor               string `json:"bg_color"`                  // ""
				BorderColor           string `json:"border_color"`              // ""
				UseImgLabel           bool   `json:"use_img_label"`             // false
				ImgLabelURIHans       string `json:"img_label_uri_hans"`        // ""
				ImgLabelURIHant       string `json:"img_label_uri_hant"`        // ""
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"` // ""
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"` // ""
			} `json:"label"`
			StartTime     int `json:"start_time"`      // 1592710727
			PaidType      int `json:"paid_type"`       // 0
			MID           int `json:"mid"`             // 188888131
			Role          int `json:"role"`            // 3
			TvVipStatus   int `json:"tv_vip_status"`   // 0
			TvVipPayType  int `json:"tv_vip_pay_type"` // 0
			TvDueDate     int `json:"tv_due_date"`     // 0
			VipRecentTime int `json:"vip_recent_time"` // 1728292199
		} `json:"vip_info"`
		Content   string `json:"content"` // ""
		PointInfo struct {
			Point       int `json:"point"`        // 0
			ExpirePoint int `json:"expire_point"` // 0
			ExpireTime  int `json:"expire_time"`  // 0
			ExpireDays  int `json:"expire_days"`  // 0
		} `json:"point_info"`
		Task struct {
			TaskItem []struct {
				TaskCode      string `json:"task_code"`      // "bonus"
				State         int    `json:"state"`          // 1
				Title         string `json:"title"`          // "大会员福利大积分"
				Icon          string `json:"icon"`           // "https://i0.hdslb.com/bfs/activity-plat/static/20220607/b66bfe4ccfd6bed05bdb54008ff5c7aa/IbtMl6R3yt.png"
				Subtitle      string `json:"subtitle"`       // "大会员/年度大会员<br /><span class="active">+100/200大积分"
				Explain       string `json:"explain"`        // "在期大会员可领取100大积分，在期年度大会员可领取200大积分，每人限领1次。如当身份为大会员时已领取过本任务大积分，则后续成为年度大会员也无法补领取差值大积分。"
				VipLimit      int    `json:"vip_limit"`      // 1
				CompleteTimes int    `json:"complete_times"` // 0
				MaxTimes      int    `json:"max_times"`      // 1
				RecallNum     int    `json:"recall_num"`     // 0
				Link          string `json:"link,omitempty"` // "https://www.bilibili.com/h5/mall/home?f_source=vip&navhide=1"
			} `json:"task_item"`
			TaskCount int  `json:"task_count"` // 9
			Signed    bool `json:"signed"`     // false
			Score     int  `json:"score"`      // 5
			ExpValue  int  `json:"exp_value"`  // -1
		} `json:"task"`
		Banner        any `json:"banner"`
		GoodsCategory []struct {
			ID    int    `json:"id"`    // 110
			Name  string `json:"name"`  // "大会员"
			State int    `json:"state"` // 2
		} `json:"goods_category"`
		GoodsSkus []struct {
			Base struct {
				Token            string   `json:"token"`   // "1126774380639117732"
				Title            string   `json:"title"`   // "百草味50元礼盒直减券"
				Picture          string   `json:"picture"` // "https://i0.hdslb.com/bfs/activity-plat/9f7fe43bb6cd1692f10b9c0a2892ea63fc80041e.jpg"
				RotationPictures []string `json:"rotation_pictures"`
				Price            struct {
					Origin    int `json:"origin"` // 188
					Promotion any `json:"promotion"`
					Sale      int `json:"sale"` // 188
				} `json:"price"`
				Inventory struct {
					AvailableNum int `json:"available_num"` // 50000
					UsedNum      int `json:"used_num"`      // 44
					SurplusNum   int `json:"surplus_num"`   // 49956
				} `json:"inventory"`
				UserType          int `json:"user_type"`           // 2
				ExchangeLimitType int `json:"exchange_limit_type"` // 2
				ExchangeLimitNum  int `json:"exchange_limit_num"`  // 5
				StartTime         int `json:"start_time"`          // 1735198320
				EndTime           int `json:"end_time"`            // 1753977599
				State             int `json:"state"`               // 2
				Priority          int `json:"priority"`            // 90
			} `json:"base"`
		} `json:"goods_skus"`
		CurrentTs       int  `json:"current_ts"`       // 1735366429
		IntegrationTask bool `json:"integration_task"` // false
	} `json:"data"`
}

// 大积分中心信息
func GetHomepageCombine(credential *Credential) (result HomepageCombineResponse, err error) {
	err = cli.Result(HomepageCombine{Credential: credential}, &result)
	return
}

// 大积分改变记录
type VIPPointList struct {
	req.Get
	*Credential

	// 改变类型
	// (0)所有类型
	// (1)获取记录
	// (2)消耗记录
	ChangeType int `api:"query:0"`

	// 分页页数
	Pn int `api:"query:0"`

	// 分页大小
	Ps int `api:"query:20"`
}

func (VIPPointList) RawURL() string {
	return "/vip_point/list"
}

func (api *VIPPointList) ReadPage(v any) (err error) {
	err = cli.Result(api, v)
	api.Pn++
	return
}

var _ PageReader = (*VIPPointList)(nil)

type VIPPointListResponse struct {
	Error
	Data struct {
		Total        int `json:"total"` // 2
		BigPointList []struct {
			Point      int    `json:"point"`       // 50
			ChangeTime int    `json:"change_time"` // 1680253998
			Remark     string `json:"remark"`      // "领取大会员卡券包权益"
			OrderNo    string `json:"order_no"`    // "nt-s-c-1018922562"
			ImageURL   string `json:"image_url"`   // ""
		} `json:"big_point_list"`
	} `json:"data"`
}

func (r VIPPointListResponse) More() bool {
	return len(r.Data.BigPointList) != 0
}

var _ Morer = (*VIPPointListResponse)(nil)

// 大积分改变记录
func GetVIPPointList(credential *Credential) (result VIPPointListResponse, err error) {
	err = cli.Result(VIPPointList{Credential: credential}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/vip/clockin.html

// 大积分签到
type TaskSign struct {
	PostCSRF
	*Credential
}

func (TaskSign) RawURL() string {
	return "https://api.bilibili.com/pgc/activity/score/task/sign"
}

type TaskSignResponse struct {
	Error
}

// 大积分签到
func PostTaskSign(credential *Credential) (result TaskSignResponse, err error) {
	err = cli.Result(TaskSign{Credential: credential}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/vip/action.html

// 大会员每日经验
type VIPEXPAdd struct {
	PostCSRF
	*Credential
}

func (VIPEXPAdd) RawURL() string {
	return "/vip/experience/add"
}

type VIPEXPAddResponse struct {
	Error
	Data struct {
		Type    int  `json:"type"`     // 0
		IsGrant bool `json:"is_grant"` // true
	} `json:"data"`
}

// 大会员每日经验
func PostVIPEXPAdd(credential *Credential) (result VIPEXPAddResponse, err error) {
	err = cli.Result(VIPEXPAdd{Credential: credential}, &result)
	return
}
