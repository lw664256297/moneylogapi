package handler

import (
	"moneylogapi/pkg/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SenResponse 返回值格式化
func SenResponse(c *gin.Context, err error, data interface{}) {
	code, message := error.DecodeErr(err)

	// 总是返回http.statusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		data:    data,
	})
}
