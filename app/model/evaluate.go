package model

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/database"
	"evaluate_backend/app/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProduct(ctx *gin.Context, db *gorm.DB, condition map[string]interface{}, page, pageSize int) (total int64, productList *[]database.Product, err error) {
	m := database.Product{}
	db = db.Model(m).Select(util.GetJsonFields(m))
	offset := util.GetOffset(page, pageSize)
	totalQuery := db
	totalQuery.Where(condition).Where("is_deleted = ?", enums.IsDeletedNo).Count(&total)
	result := db.Where(condition).Where("is_deleted = ?", enums.IsDeletedNo).Offset(offset).Limit(pageSize).Order("id desc").Find(&productList)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	return
}

func UpdateProduct(ctx *gin.Context, db *gorm.DB, condition map[string]interface{}, updateAttrs map[string]interface{}) error {
	m := database.Product{}
	result := db.Model(m).Where(condition).Updates(updateAttrs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}