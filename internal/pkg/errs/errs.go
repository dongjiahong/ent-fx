// Package errs 错误信息包
package errs

var NoRecord = "没有记录"

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	AuthCheckTokenFail    = 4001
	AuthCheckTokenTimeOut = 4002
	AuthTokenErr          = 4003
	AuthUserDisabled      = 4004

	DBError     = 5001
	CacheError  = 5002
	SmsError    = 5003
	EmailError  = 5004
	LotusError  = 5005
	ErrorNodeId = 5006
)

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "服务异常",
	InvalidParams: "请求参数错误",

	AuthCheckTokenFail:    "Token鉴权失败",
	AuthCheckTokenTimeOut: "Token已超时",
	AuthTokenErr:          "Token生成失败",
	AuthUserDisabled:      "用户被禁用",

	DBError:     "服务繁忙请稍后",
	CacheError:  "数据缓存服务异常",
	SmsError:    "短信服务异常",
	EmailError:  "邮件服务异常",
	ErrorNodeId: "错误节点id",
}

type CodeError struct {
	code int
	msg  string
}

func IsNoData(err error) bool {
	return err.Error() == NoRecord
}

func New(code int, msg string) error {
	return &CodeError{code, msg}
}

func (ce *CodeError) Error() string {
	if len(ce.msg) > 0 {
		return ce.msg
	}
	msg, ok := MsgFlags[ce.code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}

func (ce *CodeError) Code() int {
	return ce.code
}
