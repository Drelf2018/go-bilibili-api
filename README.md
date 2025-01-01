# go-bilibili-api

哔哩哔哩 API 的 go 语言实现。

接口编写参考 [BAC Document](https://socialsisteryi.github.io/bilibili-API-collect) 目前已按顺序实现至 [大会员](https://socialsisteryi.github.io/bilibili-API-collect/docs/vip) 。

## 背景知识

在使用本库前，你需要了解 [req](https://github.com/Drelf2018/req) 库的使用方法。

## 使用方法

你先别急🤭

先看本库有没有实现，如果没有实现，请告诉我，我有空会实现。

其次，本库最大的优点是自由度高，你可以按照本库中编写 `API` 的方式写你需要的接口，再用 [client.go](./client.go) 里的函数实现调用

```go
func Do[T any](api req.API) (result T, err error) {
	err = cli.Result(api, &result)
	return
}
```