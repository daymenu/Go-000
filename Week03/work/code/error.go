package code

// ErrCode 定义错误
type ErrCode int

const (
	// 1-1000 致命错误

	// ServerErrCode 服务错误
	ServerErrCode = 1 + iota

	// 2000-10000 非致命错误

	// BizErrCode 业务错误
	BizErrCode = 2000 + iota

	// 3000 日志类错误

	// LogErrCode  日志启动失败
	LogErrCode = 3000 + iota
)
