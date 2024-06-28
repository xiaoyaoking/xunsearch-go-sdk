# xunsearch-go-sdk
xunsearch golang sdk

1.兼容原生ini配置文件

2.修复字段乱序问题

3.修复返回 doc 数量错误

4.新增 SetSort 排序选项 


original author https://github.com/ninggf/xs4go
## 文档

建立索引: 参考[Xunsearch索引](http://www.xunsearch.com/doc/php/guide/index.overview)部分。

搜索: 参考[Xunsearch搜索](http://www.xunsearch.com/doc/php/guide/search.overview)部分。

## 配置文件

参考`test/demo.ini`和[Xunsearch](http://www.xunsearch.com/doc/php/guide/ini.guide)官方。

## 分词器

请自己实现如下接口：

```go
type Tokenizer interface {
    GetTokens(text string) []string
}
```

然后将其设置为分词器:

```go
index, err := xs.NewIndexer("./demo.ini")

index.SetTokenizer(yourTokenizer)
```
