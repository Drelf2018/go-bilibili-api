package api

import (
	"time"

	"github.com/Drelf2018/req"
)

// 判断是否还有新内容
type Morer interface {
	More() bool
}

// 页读器
type PageReader interface {
	// 将当前页读取至入参中并返回页码和可能产生的错误
	// 完成读取后应当改变页码 下次调用这个方法时能返回新的值
	ReadPage(any) (page int, err error)
}

// 对于需要读取多页的接口设计的迭代器
//
// 虽然迭代器要 go1.23 及以上才能用
//
// 但这不代表我这个 go1.18 的库不能先写出这样的函数来🤭
func ReadPage[V Morer](api PageReader, timer req.RetryTimer) func(yeild func(int, V) bool) {
	return func(yeild func(int, V) bool) {
		var v V
		page, err := api.ReadPage(&v)
		if err != nil {
			return
		}
		for v.More() {
			if !yeild(page, v) {
				return
			}
			d, ok := timer.NextRetry(0)
			if !ok {
				return
			}
			time.Sleep(d)
			var zero V
			page, err = api.ReadPage(&zero)
			if err != nil {
				return
			}
			v = zero
		}
	}
}
