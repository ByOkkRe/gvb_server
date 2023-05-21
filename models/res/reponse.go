package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Sucess = 0
	Error  = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(
		http.StatusOK,
		Response{
			Code: code,
			Data: data,
			Msg:  msg,
		})
}

func OK(c *gin.Context) {
	Result(Sucess, map[string]any{}, "sucess", c)
}

func OkWithData(data any, c *gin.Context) {
	Result(Sucess, data, "sucess", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(Sucess, map[string]any{}, msg, c)
}

func OkWithDetailed(data any, msg string, c *gin.Context) {
	Result(Sucess, data, msg, c)
}
func Fail(c *gin.Context) {
	Result(Error, map[string]any{}, "sucess", c)
}
func FailWithData(data any, c *gin.Context) {
	Result(Error, data, "sucess", c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}

func FailWithDetailed(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	if msg, ok := ErrorMap[code]; ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(int(code), map[string]any{}, "未知错误", c)
}
