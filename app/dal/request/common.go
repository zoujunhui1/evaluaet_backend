package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//登陆
type LoginReq struct {
	Name     string `json:"name" binding:"required"`     //用户名
	Password string `json:"password" binding:"required"` //密码
}

func (req *LoginReq) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	return nil
}
