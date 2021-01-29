package errcode

// Err 错误类型
type Err int32

// ErrSys 内部错误
const (
	ErrSysUnknown Err = -1000 + iota
)

// ErrTool 工具向错误
const (
	ErrToolUnknown Err = -700 + iota
)

// Err 业务逻辑错误
const (
	Succ Err = 1000 + iota
	ErrUnknown
)
