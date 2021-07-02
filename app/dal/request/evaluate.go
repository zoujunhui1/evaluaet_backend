package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

//产品列表
type GetProductListReq struct {
	ProductID int64 `form:"product_id" json:"product_id"`
	PageSize  int   `form:"page_size" json:"page_size"`
	Page      int   `form:"page" json:"page"`
}

func (req *GetProductListReq) Validate(ctx *gin.Context) error {
	if err := ctx.BindQuery(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	if req.ProductID < 0 {
		return errors.Errorf("product_id(%d) <= 0", req.ProductID)
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 20
	}
	return nil
}

//编辑产品
type EditProductReq struct {
	ProductID      int64  `json:"product_id" binding:"required"` //产品id
	Name           string `json:"name" binding:"required"`       //名称
	ProductType    string `json:"product_type"`                  //产品类型
	IssueTime      string `json:"issue_time"`                    //发行时间
	Denomination   string `json:"denomination"`                  //面值
	ProductVersion string `json:"product_version"`               //版别
	Weight         int32  `json:"weight"`                        //重量
	Thick          int32  `json:"thick"`                         //厚度
	Diameter       int32  `json:"diameter"`                      //直径
	Score          string `json:"score"`                         //评级分数
	IdentifyResult string `json:"identify_result"`               //鉴定结果
	Desc           string `json:"desc"`                          //备注说明
	ProductCount   int64  `json:"product_count"`                 //生成的产品数量
}

func (req *EditProductReq) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.Errorf("params ShouldBindJSON validate err(%v)", err)
	}
	return nil
}

//删除产品
type DelProductReq struct {
	ProductID int64 `json:"product_id" binding:"required"` //产品id
}

func (req *DelProductReq) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	return nil
}

//产品详情
type GetProductInfoReq struct {
	ProductID int64 `form:"product_id" json:"product_id"`
}

func (req *GetProductInfoReq) Validate(ctx *gin.Context) error {
	if err := ctx.BindQuery(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	if req.ProductID <= 0 {
		return errors.Errorf("product_id(%d) <= 0", req.ProductID)
	}
	return nil
}

//图片下载
type ImageDownloadReq struct {
	ProductID int64 `form:"product_id" json:"product_id" binding:"required"`
}

func (req *ImageDownloadReq) Validate(ctx *gin.Context) error {
	if err := ctx.BindQuery(&req); err != nil {
		return errors.Errorf("params validate err(%v)", err)
	}
	if req.ProductID <= 0 {
		return errors.Errorf("product_id(%d) <= 0", req.ProductID)
	}
	return nil
}
