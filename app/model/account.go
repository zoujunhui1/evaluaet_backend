package model

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/database"
	"evaluate_backend/app/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountUpdate(ctx *gin.Context, db *gorm.DB, condition map[string]interface{}, updateAttrs map[string]interface{}) error {
	m := database.Account{}
	result := db.Model(m).Where(condition).Updates(updateAttrs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AccountGet(ctx *gin.Context, db *gorm.DB, condition map[string]interface{}) ([]database.Account, error) {
	m := database.Account{}
	account := []database.Account{}
	db = db.Model(m).Select(util.GetJsonFields(m))
	result := db.Where(condition).Where("is_deleted = ?", enums.IsDeletedNo).Find(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}
