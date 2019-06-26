package model

//请求任务封装体
type Request struct {
	Url string
	Content string
	//url 对应的解析函数
	ParserFunc func([]byte,interface{})
	// ParserFunc func([]byte,interface{},interface{})
}

//解析结果
type ParseResult struct {
	UrlOfDownload string
	Content string
	// 解析出来的实体（例如，城市名），是任意类别（interface{}，类比 java Object）
	Items []interface{}
}