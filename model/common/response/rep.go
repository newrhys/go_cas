package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageResult struct {
	List     		interface{} `json:"list"`
	Total    		int64       `json:"total"`
	CurrentPage     int         `json:"current_page"`
	PageSize 		int         `json:"page_size"`
}

type Response struct {
	Code int			`json:"code"`
	Data interface{}	`json:"data"`
	Msg  string			`json:"msg"`
}

func Result(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Result(ctx, 200, data, msg)
}

func SuccessWithMessage(ctx *gin.Context, msg string)  {
	Result(ctx, 200, nil, msg)
}

func Fail(ctx *gin.Context, code int, msg string) {
	Result(ctx, code, nil, msg)
}

func FailWithMessage(ctx *gin.Context, msg string)  {
	Result(ctx, 400, nil, msg)
}