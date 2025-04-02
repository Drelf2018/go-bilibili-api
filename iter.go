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
	// 读取当前页
	// 读取后改变页码以便下次调用时能返回新的值
	ReadPage(any) error
}

// 对于需要读取多页的接口设计的迭代器（旧）
func ReadPageOld[V Morer](api PageReader, ticker req.RetryTicker) func(yield func(int, V) bool) {
	return func(yield func(int, V) bool) {
		var v V
		err := api.ReadPage(&v)
		if err != nil {
			return
		}
		for v.More() {
			if !yield(0, v) {
				return
			}
			d, ok := ticker.NextRetry(0)
			if !ok {
				return
			}
			time.Sleep(d)
			var zero V
			err = api.ReadPage(&zero)
			if err != nil {
				return
			}
			v = zero
		}
	}
}

// 对于需要读取多页的接口设计的迭代器
//
// 虽然迭代器要 go1.23 及以上才能用
//
// 但这不代表我这个 go1.18 的库不能先写出这样的函数来🤭
func ReadPage[V Morer](api PageReader, ticker req.RetryTicker) func(func(V) bool) {
	return func(yield func(V) bool) {
		for i := 0; ; i++ {
			var zero V
			err := api.ReadPage(&zero)
			d, ok := ticker.NextRetry(i)
			if err != nil || !yield(zero) || !zero.More() || !ok {
				return
			}
			time.Sleep(d)
		}
	}
}
