package api

import (
	"errors"
	"math/rand"
	"time"
)

var ErrNoMorePage = errors.New("api: no more page")

// 页读器
type PageReader[V any] interface {
	// 读取当前页
	// 读取后改变页码以便下次调用时能返回新的值
	// 没有新内容返回错误
	// 	return v, ErrNoMorePage
	ReadPage() (V, error)
}

// 对于需要读取多页的接口设计的迭代器
//
// 虽然迭代器要 go1.23 及以上才能用
//
// 但这不代表我这个 go1.18 的库不能先写出这样的函数来🤭
func ReadPage[V any](api PageReader[V], interval ...time.Duration) func(func(V) bool) {
	return func(yield func(V) bool) {
		for v, err := api.ReadPage(); err == nil && yield(v); v, err = api.ReadPage() {
			var duration time.Duration
			for i, d := range interval {
				if d <= 0 {
					continue
				}
				if i == 0 {
					duration = d
				} else {
					duration += time.Duration(rand.Int63n(int64(d)) / (1 << i))
				}
			}
			time.Sleep(duration)
		}
	}
}
