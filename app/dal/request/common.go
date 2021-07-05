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

//增加枚举
type AddEnumReq struct {
	EnumID   int    `json:"enum_id"`
	EnumName string `json:"enum_name"`
}

func (req *AddEnumReq) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	return nil
}

//删除枚举
type DelEnumReq struct {
	EnumID int `json:"enum_id" binding:"required"`
}

func (req *DelEnumReq) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	return nil
}

//获取枚举
type GetEnumReq struct {
	EnumID   int `form:"enum_id" json:"enum_id"`
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}

func (req *GetEnumReq) Validate(ctx *gin.Context) error {
	if err := ctx.BindQuery(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	return nil

}
