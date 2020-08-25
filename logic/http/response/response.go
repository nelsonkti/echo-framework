package response

// 请求返回值
const (
    CodeSuccess   = 0 // 成功返回
    CodeFail      = 1
    CodeFailRetry = 2 //需要改参数重试
)

// Response 用户响应数据
type RespData struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

//请求成功，返回
func Success(data interface{}) RespData {
    return RespData{
        Code: CodeSuccess,
        Msg:  "成功",
        Data: data,
    }
}

//请求成功，返回
func Fail(msg string) RespData {
    return RespData{
        Code: CodeFail,
        Msg:  msg,
        Data: nil,
    }
}

//失败
func FailCode(msg string, code int) RespData {
    return RespData{
        Code: code,
        Msg:  msg,
        Data: nil,
    }
}
