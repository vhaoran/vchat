package ykit

/*--auth: whr  date:2019/12/511:43--------------------------
 ####请勿擅改此功能代码####
 用途：
 统一的服务返回数据接口，用于标准化输出
--------------------------------------- */
type (
	Result struct {
		//正常返回值为200，
		//小于0为错误，大于0的非200值为有警告
		Code int `json:"code"`
		//文本消息，错误内容或警告内容，可执行成功
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)

const (
	OK      = 200
	ERR     = -1
	WARNING = 1000
)

func RWarn(msg string, data interface{}) *Result {
	return &Result{
		Code: WARNING,
		Msg:  msg,
		Data: data,
	}
}

func ROK(data interface{}, info ...string) *Result {
	return ResultOK(data, info...)
}

func ResultOK(data interface{}, info ...string) *Result {
	s := "执行成功"
	if len(info) > 0 {
		s += "," + info[0]
	}

	return &Result{
		Code: OK,
		Msg:  s,
		Data: data,
	}
}

func RErr(err error) *Result {
	return ResultError(err)
}

func ResultError(err error) *Result {
	return &Result{
		Code: ERR,
		Msg:  "操作错误" + err.Error(),
		Data: err.Error(),
	}
}

func ResultWarn(msg string, data interface{}) *Result {
	return ResultWarning(msg, data)
}

func ResultWarning(warnMsg string, data interface{}) *Result {
	return &Result{
		Code: WARNING,
		Msg:  warnMsg,
		Data: data,
	}
}
