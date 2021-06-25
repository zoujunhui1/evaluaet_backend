package model

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/database"
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProduct(ctx *gin.Context, db *gorm.DB, req *request.GetProductListReq) (total int64, productList *[]database.Product, err error) {
	m := database.Product{}
	db = db.Model(m).Select(util.GetJsonFields(m))
	condition := make(map[string]interface{})
	offset := util.GetOffset(req.Page, req.PageSize)
	if req.ProductID > 0 {
		condition["product_id"] = req.ProductID
	}
	totalQuery := db
	totalQuery.Where(condition).Where("is_deleted = ?", enums.IsDeletedNo).Count(&total)
	result := db.Where(condition).Where("is_deleted = ?", enums.IsDeletedNo).Offset(offset).Limit(req.PageSize).Order("id desc").Find(&productList)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	return
}
