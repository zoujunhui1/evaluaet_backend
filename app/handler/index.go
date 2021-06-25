package handler

import (
	"evaluate_backend/app/const/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ResponseBody 响应Body
type ResponseBody struct {
	Status  int32       `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Fail(ctx *gin.Context, code int32) {
	status := enums.StatusCode(code)
	body := &ResponseBody{
		Status:  status.Code(),
		Message: status.Message(),
		Data:    &map[string]string{},
	}
	ctx.Set("response", body)
	ctx.JSON(http.StatusOK, body)
}

func Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = &map[string]string{}
	}

	status := enums.StatusCode(enums.StatusOK)
	body := &ResponseBody{
		Status:  status.Code(),
		Message: status.Message(),
		Data:    data,
	}
	ctx.Set("response", body)
	ctx.JSON(http.StatusOK, body)
}
