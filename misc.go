package api

import "github.com/Drelf2018/req"

// https://socialsisteryi.github.io/bilibili-API-collect/docs/misc/time_stamp.html

// 获取当前时间戳
type Now struct {
	req.Get
}

func (Now) RawURL() string {
	return "/report/click/now"
}

type NowResponse struct {
	Error
	Data struct {
		Now int `json:"now"` // 1735223842
	} `json:"data"`
}

// 获取当前时间戳
func GetNow() (result NowResponse, err error) {
	err = cli.Result(Now{}, &result)
	return
}

// 获取适用于 RTC 的时间戳
type Timestamp struct {
	req.Get
}

func (Timestamp) RawURL() string {
	return "https://api.live.bilibili.com/xlive/open-interface/v1/rtc/getTimestamp"
}

type TimestampResponse struct {
	Error
	Data struct {
		Timestamp int   `json:"timestamp"` // 1735223942
		Microtime int64 `json:"microtime"` // 1735223942970
	} `json:"data"`
}

// 获取适用于 RTC 的时间戳
func GetTimestamp() (result TimestampResponse, err error) {
	err = cli.Result(Timestamp{}, &result)
	return
}

// https://socialsisteryi.github.io/bilibili-API-collect/docs/misc/buvid3_4.html

// 接口获取 buvid3 / buvid4
type SPI struct {
	req.Get
}

func (SPI) RawURL() string {
	return "/frontend/finger/spi"
}

type SPIResponse struct {
	Error
	Data struct {
		B3 string `json:"b_3"` // "E225B4F4-615F-F3D4-6AD3-EA4CDD9B3D5389718infoc"
		B4 string `json:"b_4"` // "160D49C2-BDF9-6152-15B5-3607E9641BF789718-024122614-kzVGhgiChnkj56LAFFA0BA=="
	} `json:"data"`
}

// 接口获取 buvid3 / buvid4
func GetSPI() (result SPIResponse, err error) {
	err = cli.Result(SPI{}, &result)
	return
}

// 接口获取 buvid3
func GetBuvid3() string {
	r, err := GetSPI()
	if err != nil {
		return ""
	}
	return r.Data.B3
}
