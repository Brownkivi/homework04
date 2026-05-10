package exception

// BizError 自定义业务错误
type BizError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// NewBizError 创建业务错误
func NewBizError(code int, msg string) *BizError {
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

// 实现 error 接口
func (e *BizError) Error() string {
	return e.Msg
}
