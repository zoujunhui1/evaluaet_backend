package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"mime/multipart"
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

//退出
type LogoutReq struct {
	ID    int64  `json:"id" binding:"required"`
	Token string `json:"token"`
}

func (req *LogoutReq) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	return nil
}

//图片上传
type ImageUploadReq struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (req *ImageUploadReq) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBind(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	return nil
}
