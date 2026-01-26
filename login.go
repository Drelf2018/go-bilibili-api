package bilibili

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/Drelf2018/req"

	_ "embed"
)

// https://socialsisteryi.github.io/bilibili-API-collect/docs/login/login_action/QR.html

// 申请二维码
type Generate struct {
	req.Get
}

func (Generate) RawURL() string {
	return "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"
}

type GenerateResponse struct {
	Error
	Data struct {
		URL       string `json:"url"`        // https://account.bilibili.com/h5/account-h5/auth/scan-web?navhide=1&callback=close&qrcode_key=52360fad71935c52ca33f4f24fd18e07&from=
		QRCodeKey string `json:"qrcode_key"` // 52360fad71935c52ca33f4f24fd18e07
	} `json:"data"`
}

// 申请二维码
func GetGenerate(ctx context.Context) (GenerateResponse, error) {
	return Do[GenerateResponse](ctx, Generate{})
}

// 扫码登录
//
// 登录成功后会自动将 Cookie 值写入该结构体的 *Credential 字段中
type Poll struct {
	req.Get
	http.CookieJar

	// 先前生成的二维码密钥
	QRCodeKey string `req:"query:qrcode_key"`
}

func (Poll) RawURL() string {
	return "https://passport.bilibili.com/x/passport-login/web/qrcode/poll"
}

type PollResponse struct {
	Error
	Data struct {
		Error
		URL          string `json:"url"`
		RefreshToken string `json:"refresh_token"`
		Timestamp    int    `json:"timestamp"`
	} `json:"data"`
}

func (r PollResponse) Unwrap() error {
	err := r.Error.Unwrap()
	if err != nil {
		return err
	}
	return r.Data.Error.Unwrap()
}

// 登录成功后会自动将 Cookie 值写入 credential 变量中
func GetPoll(ctx context.Context, jar http.CookieJar, qrcodeKey string) (PollResponse, error) {
	return Do[PollResponse](ctx, Poll{QRCodeKey: qrcodeKey, CookieJar: jar})
}

// 登录成功后会自动将 Cookie 值写入返回值 *Credential 中
func GetCredential(ctx context.Context, qrcodeKey string) (*Credential, error) {
	cred := &Credential{}
	r, err := GetPoll(ctx, cred, qrcodeKey)
	if err != nil {
		return nil, err
	}
	cred.RefreshToken = RefreshToken(r.Data.RefreshToken)
	return cred, nil
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/login/login_info.html

// 导航栏用户信息
type Nav struct {
	req.Get
	http.CookieJar
}

func (Nav) RawURL() string {
	return "/web-interface/nav"
}

// 完整返回值结构
type NavResponse struct {
	Error
	Data struct {
		IsLogin       bool   `json:"isLogin"` // false
		EmailVerified int    `json:"email_verified"`
		Face          string `json:"face"`
		FaceNft       int    `json:"face_nft"`
		FaceNftType   int    `json:"face_nft_type"`
		LevelInfo     struct {
			CurrentLevel int    `json:"current_level"`
			CurrentMin   int    `json:"current_min"`
			CurrentExp   int    `json:"current_exp"`
			NextExp      string `json:"next_exp"`
		} `json:"level_info"`
		MID            int     `json:"mid"`
		MobileVerified int     `json:"mobile_verified"`
		Money          float64 `json:"money"`
		Moral          int     `json:"moral"`
		Official       struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"official"`
		OfficialVerify struct {
			Type int    `json:"type"`
			Desc string `json:"desc"`
		} `json:"officialVerify"`
		Pendant struct {
			PID               int    `json:"pid"`
			Name              string `json:"name"`
			Image             string `json:"image"`
			Expire            int    `json:"expire"`
			ImageEnhance      string `json:"image_enhance"`
			ImageEnhanceFrame string `json:"image_enhance_frame"`
			NPID              int    `json:"n_pid"`
		} `json:"pendant"`
		Scores       int    `json:"scores"`
		Uname        string `json:"uname"`
		VipDueDate   int64  `json:"vipDueDate"`
		VipStatus    int    `json:"vipStatus"`
		VipType      int    `json:"vipType"`
		VipPayType   int    `json:"vip_pay_type"`
		VipThemeType int    `json:"vip_theme_type"`
		VipLabel     struct {
			Path                  string `json:"path"`
			Text                  string `json:"text"`
			LabelTheme            string `json:"label_theme"`
			TextColor             string `json:"text_color"`
			BgStyle               int    `json:"bg_style"`
			BgColor               string `json:"bg_color"`
			BorderColor           string `json:"border_color"`
			UseImgLabel           bool   `json:"use_img_label"`
			ImgLabelURIHans       string `json:"img_label_uri_hans"`
			ImgLabelURIHant       string `json:"img_label_uri_hant"`
			ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
			ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
		} `json:"vip_label"`
		VipAvatarSubscript int    `json:"vip_avatar_subscript"`
		VipNicknameColor   string `json:"vip_nickname_color"`
		Vip                struct {
			Type       int   `json:"type"`
			Status     int   `json:"status"`
			DueDate    int64 `json:"due_date"`
			VipPayType int   `json:"vip_pay_type"`
			ThemeType  int   `json:"theme_type"`
			Label      struct {
				Path                  string `json:"path"`
				Text                  string `json:"text"`
				LabelTheme            string `json:"label_theme"`
				TextColor             string `json:"text_color"`
				BgStyle               int    `json:"bg_style"`
				BgColor               string `json:"bg_color"`
				BorderColor           string `json:"border_color"`
				UseImgLabel           bool   `json:"use_img_label"`
				ImgLabelURIHans       string `json:"img_label_uri_hans"`
				ImgLabelURIHant       string `json:"img_label_uri_hant"`
				ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
				ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
			} `json:"label"`
			AvatarSubscript    int    `json:"avatar_subscript"`
			NicknameColor      string `json:"nickname_color"`
			Role               int    `json:"role"`
			AvatarSubscriptURL string `json:"avatar_subscript_url"`
			TvVipStatus        int    `json:"tv_vip_status"`
			TvVipPayType       int    `json:"tv_vip_pay_type"`
			TvDueDate          int    `json:"tv_due_date"`
			AvatarIcon         struct {
				IconType     int `json:"icon_type"`
				IconResource struct {
				} `json:"icon_resource"`
			} `json:"avatar_icon"`
		} `json:"vip"`
		Wallet struct {
			MID           int `json:"mid"`
			BcoinBalance  int `json:"bcoin_balance"`
			CouponBalance int `json:"coupon_balance"`
			CouponDueTime int `json:"coupon_due_time"`
		} `json:"wallet"`
		HasShop        bool   `json:"has_shop"`
		ShopURL        string `json:"shop_url"`
		AllowanceCount int    `json:"allowance_count"`
		AnswerStatus   int    `json:"answer_status"`
		IsSeniorMember int    `json:"is_senior_member"`
		WbiImg         struct {
			ImgURL string `json:"img_url"` // https://i0.hdslb.com/bfs/wbi/7cd084941338484aae1ad9425b84077c.png
			SubURL string `json:"sub_url"` // https://i0.hdslb.com/bfs/wbi/4932caff0ff746eab6f01bf08b70ac45.png
		} `json:"wbi_img"`
		IsJury     bool        `json:"is_jury"`
		NameRender interface{} `json:"name_render"`
	} `json:"data"`
}

// 未登录返回值结构
type NavResponseSimple struct {
	Error
	Data struct {
		IsLogin bool `json:"isLogin"` // false
		WbiImg  struct {
			ImgURL string `json:"img_url"` // https://i0.hdslb.com/bfs/wbi/7cd084941338484aae1ad9425b84077c.png
			SubURL string `json:"sub_url"` // https://i0.hdslb.com/bfs/wbi/4932caff0ff746eab6f01bf08b70ac45.png
		} `json:"wbi_img"`
	} `json:"data"`
}

// 导航栏用户信息
func GetNav(ctx context.Context, jar http.CookieJar) (NavResponseSimple, error) {
	return Do[NavResponseSimple](ctx, Nav{CookieJar: jar})
}

// 登录用户状态数
type NavStat struct {
	req.Get
	http.CookieJar
}

func (NavStat) RawURL() string {
	return "/web-interface/nav/stat"
}

type NavStatResponse struct {
	Error
	Data struct {
		Following    int `json:"following"`
		Follower     int `json:"follower"`
		DynamicCount int `json:"dynamic_count"`
	} `json:"data"`
}

// 登录用户状态数
func GetNavStat(ctx context.Context, jar http.CookieJar) (NavStatResponse, error) {
	return Do[NavStatResponse](ctx, NavStat{CookieJar: jar})
}

// 获取硬币数
type Coin struct {
	req.Get
	http.CookieJar
}

func (Coin) RawURL() string {
	return "https://account.bilibili.com/site/getCoin"
}

type CoinResponse struct {
	Error
	Status bool `json:"status"`
	Data   struct {
		Money float64 `json:"money"`
	} `json:"data"`
}

// 获取硬币数
func GetCoin(ctx context.Context, jar http.CookieJar) (CoinResponse, error) {
	return Do[CoinResponse](ctx, Coin{CookieJar: jar})
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/login/member_center.html

// 获取我的信息
type Account struct {
	req.Get
	http.CookieJar
}

func (Account) RawURL() string {
	return "/member/web/account"
}

type AccountResponse struct {
	Error
	Data struct {
		MID      int    `json:"mid"`
		Uname    string `json:"uname"`
		UserID   string `json:"userid"`
		Sign     string `json:"sign"`
		Birthday string `json:"birthday"`
		Sex      string `json:"sex"`
		NickFree bool   `json:"nick_free"`
		Rank     string `json:"rank"`
	} `json:"data"`
}

// 获取我的信息
func GetAccount(ctx context.Context, jar http.CookieJar) (AccountResponse, error) {
	return Do[AccountResponse](ctx, Account{CookieJar: jar})
}

// 查询每日奖励状态
type Reward struct {
	req.Get
	http.CookieJar
}

func (Reward) RawURL() string {
	return "/member/web/exp/reward"
}

type RewardResponse struct {
	Error
	Data struct {
		Login        bool `json:"login"`
		Watch        bool `json:"watch"`
		Coins        int  `json:"coins"`
		Share        bool `json:"share"`
		Email        bool `json:"email"`
		Tel          bool `json:"tel"`
		SafeQuestion bool `json:"safe_question"`
		IdentifyCard bool `json:"identify_card"`
	} `json:"data"`
}

// 查询每日奖励状态
func GetReward(ctx context.Context, jar http.CookieJar) (RewardResponse, error) {
	return Do[RewardResponse](ctx, Reward{CookieJar: jar})
}

// 查询大会员状态
type WebUserInfo struct {
	req.Get
	http.CookieJar
}

func (WebUserInfo) RawURL() string {
	return "/vip/web/user/info"
}

type WebUserInfoResponse struct {
	Error
	Data struct {
		VipPayType         int    `json:"vip_pay_type"`
		VipIsAnnual        bool   `json:"vip_is_annual"`
		VipKeepTime        int    `json:"vip_keep_time"`
		MID                int    `json:"mid"`
		VipIsValid         bool   `json:"vip_is_valid"`
		AvatarSubscriptURL string `json:"avatar_subscript_url"`
		IsAutoRenew        bool   `json:"is_auto_renew"`
		NicknameColor      string `json:"nickname_color"`
		IsNewUser          bool   `json:"is_new_user"`
		VipDueDate         int64  `json:"vip_due_date"`
		VipIsOverdue       bool   `json:"vip_is_overdue"`
		VipType            int    `json:"vip_type"`
		Label              struct {
			Text                  string `json:"text"`
			LabelTheme            string `json:"label_theme"`
			ImgLabelURIHans       string `json:"img_label_uri_hans"`
			ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
			TextColor             string `json:"text_color"`
			BgStyle               int    `json:"bg_style"`
			BorderColor           string `json:"border_color"`
			BgColor               string `json:"bg_color"`
			UseImgLabel           bool   `json:"use_img_label"`
			ImgLabelURIHant       string `json:"img_label_uri_hant"`
			ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
		} `json:"label"`
		VipStatus       int  `json:"vip_status"`
		AvatarSubscript int  `json:"avatar_subscript"`
		VipExpireDays   int  `json:"vip_expire_days"`
		ThemeType       int  `json:"theme_type"`
		TipMaterial     any  `json:"tip_material"`
		VipIsMonth      bool `json:"vip_is_month"`
	} `json:"data"`
}

// 查询大会员状态
func GetWebUserInfo(ctx context.Context, jar http.CookieJar) (WebUserInfoResponse, error) {
	return Do[WebUserInfoResponse](ctx, WebUserInfo{CookieJar: jar})
}

// 查询账号安全情况
type SiteUserInfo struct {
	req.Get
	http.CookieJar
}

func (SiteUserInfo) RawURL() string {
	return "https://passport.bilibili.com/web/site/user/info"
}

type SiteUserInfoResponse struct {
	Error
	Data struct {
		AccountInfo struct {
			HideTel           string `json:"hide_tel"`
			HideMail          string `json:"hide_mail"`
			BindTel           bool   `json:"bind_tel"`
			BindMail          bool   `json:"bind_mail"`
			TelVerify         bool   `json:"tel_verify"`
			MailVerify        bool   `json:"mail_verify"`
			UnneededCheck     bool   `json:"unneeded_check"`
			RealnameCertified bool   `json:"realname_certified"`
		} `json:"account_info"`
		AccountSafe struct {
			Score    int  `json:"Score"`
			ScoreNew int  `json:"score_new"`
			PwdLevel int  `json:"pwd_level"`
			Security bool `json:"security"`
		} `json:"account_safe"`
		AccountSns struct {
			WeiboBind  int `json:"weibo_bind"`
			QqBind     int `json:"qq_bind"`
			WechatBind int `json:"wechat_bind"`
		} `json:"account_sns"`
		AccountOther struct {
			SkipVerify bool `json:"skipVerify"`
		} `json:"account_other"`
	} `json:"data"`
}

// 查询账号安全情况
func GetSiteUserInfo(ctx context.Context, jar http.CookieJar) (SiteUserInfoResponse, error) {
	return Do[SiteUserInfoResponse](ctx, SiteUserInfo{CookieJar: jar})
}

// 查询账号实名认证状态
type RealnameStatus struct {
	req.Get
	http.CookieJar
}

func (RealnameStatus) RawURL() string {
	return "/member/realname/status"
}

type RealnameStatusResponse struct {
	Error
	Data struct {
		Status int `json:"status"`
	} `json:"data"`
}

// 查询账号实名认证状态
func GetRealnameStatus(ctx context.Context, jar http.CookieJar) (RealnameStatusResponse, error) {
	return Do[RealnameStatusResponse](ctx, RealnameStatus{CookieJar: jar})
}

// 查询实名认证详细信息
type ApplyStatus struct {
	req.Get
	http.CookieJar
}

func (ApplyStatus) RawURL() string {
	return "/member/realname/apply/status"
}

type ApplyStatusResponse struct {
	Error
	Data struct {
		Status   int    `json:"status"`
		Remark   string `json:"remark"`
		Realname string `json:"realname"`
		Card     string `json:"card"`
		CardType int    `json:"card_type"`
	} `json:"data"`
}

// 查询实名认证详细信息
func GetApplyStatus(ctx context.Context, jar http.CookieJar) (ApplyStatusResponse, error) {
	return Do[ApplyStatusResponse](ctx, ApplyStatus{CookieJar: jar})
}

// 查询硬币变化情况
type CoinLog struct {
	req.Get
	http.CookieJar
}

func (CoinLog) RawURL() string {
	return "/member/web/coin/log"
}

type CoinLogResponse struct {
	Error
	Data struct {
		List []struct {
			Time   string `json:"time"`
			Delta  int    `json:"delta"`
			Reason string `json:"reason"`
		} `json:"list"`
		Count int `json:"count"`
	} `json:"data"`
}

// 查询硬币变化情况
func GetCoinLog(ctx context.Context, jar http.CookieJar) (CoinLogResponse, error) {
	return Do[CoinLogResponse](ctx, CoinLog{CookieJar: jar})
}

// 修改个人签名
type SignUpdate struct {
	PostCSRF
	http.CookieJar

	// 要设置的签名内容 可为空
	UserSign string `req:"body"`
}

func (SignUpdate) RawURL() string {
	return "/member/web/sign/update"
}

type SignUpdateResponse struct {
	Error
}

// 修改个人签名
func PostSignUpdate(ctx context.Context, jar http.CookieJar, sign string) (SignUpdateResponse, error) {
	return Do[SignUpdateResponse](ctx, SignUpdate{UserSign: sign, CookieJar: jar})
}

// 最近一周的经验记录
type EXPLog struct {
	req.Get
	http.CookieJar

	JSONP       string `req:"query" default:"jsonp"`
	WebLocation string `req:"query" default:"333.33"`
}

func (EXPLog) RawURL() string {
	return "/member/web/exp/log"
}

type EXPLogResponse struct {
	Error
	Data struct {
		List []struct {
			Delta  int    `json:"delta"`
			Time   string `json:"time"`
			Reason string `json:"reason"`
		} `json:"list"`
		Count int `json:"count"`
	} `json:"data"`
}

// 最近一周的经验记录
func GetEXPLog(ctx context.Context, jar http.CookieJar) (EXPLogResponse, error) {
	return Do[EXPLogResponse](ctx, EXPLog{CookieJar: jar})
}

// 最近一周的节操记录
type MoralLog struct {
	req.Get
	http.CookieJar

	JSONP       string `req:"query" default:"jsonp"`
	WebLocation string `req:"query" default:"333.33"`
}

func (MoralLog) RawURL() string {
	return "/member/web/moral/log"
}

type MoralLogResponse struct {
	Error
	Data struct {
		Moral int           `json:"moral"`
		List  []interface{} `json:"list"`
		Count int           `json:"count"`
	} `json:"data"`
}

// 最近一周的节操记录
func GetMoralLog(ctx context.Context, jar http.CookieJar) (MoralLogResponse, error) {
	return Do[MoralLogResponse](ctx, MoralLog{CookieJar: jar})
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/login/login_notice.html

// 查询登录记录
type LoginNotice struct {
	req.Get
	http.CookieJar

	// 用户 UID（一般是自己的）
	MID int `req:"query"`

	// 设备虚拟 ID 可为空
	Buvid3 string `req:"query"`
}

func (LoginNotice) RawURL() string {
	return "/safecenter/login_notice"
}

type LoginNoticeResponse struct {
	Error
	Data struct {
		MID        int    `json:"mid"`
		DeviceName string `json:"device_name"`
		LoginType  string `json:"login_type"`
		LoginTime  string `json:"login_time"`
		Location   string `json:"location"`
		IP         string `json:"ip"`
	} `json:"data"`
}

// 查询登录记录
func GetLoginNotice(ctx context.Context, jar http.CookieJar, uid int) (LoginNoticeResponse, error) {
	return Do[LoginNoticeResponse](ctx, LoginNotice{MID: uid, CookieJar: jar})
}

// 最近一周的登录情况
type LoginLog struct {
	req.Get
	http.CookieJar

	JSONP       string `req:"query" default:"jsonp"`
	WebLocation string `req:"query" default:"333.33"`
}

func (LoginLog) RawURL() string {
	return "/member/web/login/log"
}

type LoginLogResponse struct {
	Error
	Data struct {
		Count int `json:"count"`
		List  []struct {
			IP     string `json:"ip"`
			Time   int    `json:"time"`
			TimeAt string `json:"time_at"`
			Status bool   `json:"status"`
			Type   int    `json:"type"`
			Geo    string `json:"geo"`
		} `json:"list"`
	} `json:"data"`
}

// 最近一周的登录情况
func GetLoginLog(ctx context.Context, jar http.CookieJar) (LoginLogResponse, error) {
	return Do[LoginLogResponse](ctx, LoginLog{CookieJar: jar})
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/login/cookie_refresh.html

// 检查是否需要刷新
type CookieInfo struct {
	req.Get
	http.CookieJar
}

func (CookieInfo) RawURL() string {
	return "https://passport.bilibili.com/x/passport-login/web/cookie/info"
}

type CookieInfoResponse struct {
	Error
	Data struct {
		Refresh   bool  `json:"refresh"`   // true
		Timestamp int64 `json:"timestamp"` // 1734963138171
	} `json:"data"`
}

// 检查是否需要刷新
func GetCookieInfo(ctx context.Context, jar http.CookieJar) (CookieInfoResponse, error) {
	return Do[CookieInfoResponse](ctx, CookieInfo{CookieJar: jar})
}

//go:embed pubkey.pem
var PublicKeyPEM []byte

// 生成 CorrespondPath
//
// ts 为当前毫秒级时间戳
//
// 代码由 https://socialsisteryi.github.io/bilibili-API-collect/docs/login/cookie_refresh.html 提供
func GetCorrespondPath(ts int64) (string, error) {
	pubKeyBlock, _ := pem.Decode(PublicKeyPEM)

	pubInterface, parseErr := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	if parseErr != nil {
		return "", parseErr
	}
	pub := pubInterface.(*rsa.PublicKey)

	encryptedData, encryptErr := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, []byte(fmt.Sprintf("refresh_%d", ts)), nil)
	if encryptErr != nil {
		return "", encryptErr
	}
	return hex.EncodeToString(encryptedData), nil
}

// 获取 refresh_csrf
type Correspond struct {
	req.Get
	http.CookieJar

	// 通过 GetCorrespondPath 获取
	CorrespondPath string
}

func (c Correspond) RawURL() string {
	return fmt.Sprintf("https://www.bilibili.com/correspond/1/%s", c.CorrespondPath)
}

var refreshCSRFPattern = regexp.MustCompile(`<div id="1-name">(.*?)</div>`)

var ErrRefreshCSRFNotExist = errors.New("bilibili: refresh_csrf does not exist")

// 获取 refresh_csrf
func GetRefreshCSRF(ctx context.Context, jar http.CookieJar) (string, error) {
	path, err := GetCorrespondPath(time.Now().UnixMilli())
	if err != nil {
		return "", err
	}
	s, err := Session.TextWithContext(ctx, Correspond{CorrespondPath: path, CookieJar: jar})
	if err != nil {
		return "", err
	}
	r := refreshCSRFPattern.FindStringSubmatch(s)
	if len(r) < 2 {
		return "", ErrRefreshCSRFNotExist
	}
	return r[1], nil
}

// 刷新 Cookie
type CookieRefresh struct {
	PostCSRF
	http.CookieJar

	// 实时刷新口令
	// 通过 GetRefreshCSRF 获得
	RefreshCSRF string `req:"body:refresh_csrf"`

	// 访问来源 默认值 main_web
	Source string `req:"body" default:"main_web"`

	// 持久化刷新口令
	RefreshToken string `req:"body"`
}

func (CookieRefresh) RawURL() string {
	return "https://passport.bilibili.com/x/passport-login/web/cookie/refresh"
}

type CookieRefreshResponse struct {
	Error
	Data struct {
		Status       int    `json:"status"`        // 0
		Message      string `json:"message"`       // ""
		RefreshToken string `json:"refresh_token"` // "xxx"
	} `json:"data"`
}

// 刷新 Cookie
func PostCookieRefresh(ctx context.Context, jar http.CookieJar, token string) (result CookieRefreshResponse, err error) {
	key, err := GetRefreshCSRF(ctx, jar)
	if err != nil {
		return
	}
	return Do[CookieRefreshResponse](ctx, CookieRefresh{RefreshCSRF: key, RefreshToken: token, CookieJar: jar})
}

// 确认更新
type ConfirmRefresh struct {
	PostCSRF
	http.CookieJar

	// 旧的持久化刷新口令
	RefreshToken string `req:"body"`
}

func (ConfirmRefresh) RawURL() string {
	return "https://passport.bilibili.com/x/passport-login/web/confirm/refresh"
}

type ConfirmRefreshResponse struct {
	Error
}

// 确认更新
func PostConfirmRefresh(ctx context.Context, jar http.CookieJar, token string) (err error) {
	_, err = Do[ConfirmRefreshResponse](ctx, ConfirmRefresh{RefreshToken: token, CookieJar: jar})
	return
}
