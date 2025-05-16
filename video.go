package api

import "github.com/Drelf2018/req"

// https://socialsisteryi.github.io/bilibili-API-collect/docs/video/collection.html

type Archive struct {
	AID      int64  `json:"aid" gorm:"column:aid"` // 114410580345016
	Title    string `json:"title"`                 // "【直播回放】你给我进来！ 2025年04月27日19点场"
	Pubdate  int    `json:"pubdate"`               // 1745780073
	Ctime    int    `json:"ctime"`                 // 1745780073
	State    int    `json:"state"`                 // 0
	Pic      string `json:"pic"`                   // "http://i2.hdslb.com/bfs/archive/f75be7f33a2b8ef29e2f80b542fc787e47695658.jpg"
	Duration int    `json:"duration"`              // 10256
	Stat     struct {
		View int `json:"view"` // 251
	} `json:"stat" gorm:"serializer:json"`
	BVID             string `json:"bvid" gorm:"column:bvid"` // "BV1ZSL6zNEoB"
	UgcPay           int    `json:"ugc_pay"`                 // 0
	InteractiveVideo bool   `json:"interactive_video"`       // false
	EnableVt         int    `json:"enable_vt"`               // 0
	VtDisplay        string `json:"vt_display"`              // ""
	PlaybackPosition int    `json:"playback_position"`       // 0
}

func (a Archive) String() string {
	return a.Title
}

// 获取指定系列视频
type SeriesArchives struct {
	req.Get

	// 用户 mid
	MID int `api:"query"`

	// 系列 ID
	SeriesID int `api:"query"`

	// 作用尚不明确
	// 默认为 true
	OnlyNormal bool `api:"query"`

	// 排序方式
	// 默认排序 desc
	// 升序排序 asc
	Sort string `api:"query"`

	// 页码
	// 默认为 1
	Pn int `api:"query:1"`

	// 每页项数
	// 默认为 30
	Ps int `api:"query:30"`

	// 当前用户 UID
	CurrentMID int `api:"query"`
}

func (SeriesArchives) RawURL() string {
	return "/series/archives"
}

func (api *SeriesArchives) ReadPage() (v SeriesArchivesResponse, err error) {
	if api.Pn == 0 {
		api.Pn = 1
	}
	err = cli.Result(api, &v)
	if err != nil {
		return
	}
	if len(v.Data.Archives) == 0 {
		err = ErrNoMorePage
		return
	}
	api.Pn++
	return
}

var _ PageReader[SeriesArchivesResponse] = (*SeriesArchives)(nil)

type SeriesArchivesResponse struct {
	Error
	Data struct {
		AIDs []int64 `json:"aids"`
		Page struct {
			Num   int `json:"num"`   // 1
			Size  int `json:"size"`  // 20
			Total int `json:"total"` // 777
		} `json:"page"`
		Archives []Archive `json:"archives"`
	} `json:"data"`
}

// 获取指定系列视频
func GetSeriesArchives(uid int, seriesID int) (result SeriesArchivesResponse, err error) {
	err = cli.Result(SeriesArchives{MID: uid, SeriesID: seriesID}, &result)
	return
}
