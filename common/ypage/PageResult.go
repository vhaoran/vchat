package ypage

type PageResult struct {
	PageBean
	Code int `json:"code"`
	//文本消息，错误内容或警告内容，可执行成功
	Msg string `json:"msg"`
}


