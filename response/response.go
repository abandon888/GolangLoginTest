package response

import "github.com/gin-gonic/gin"

// Response 响应请求
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Success 响应成功
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, 200, 200, data, msg)
}

// Fail 响应失败
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, 200, 400, data, msg)
}
