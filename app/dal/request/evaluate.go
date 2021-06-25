package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

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
