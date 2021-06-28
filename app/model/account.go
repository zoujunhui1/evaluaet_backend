package model

import (
	"context"
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/database"
	"evaluate_backend/app/provider"
	"evaluate_backend/app/util"
)

func AccountUpdate(ctx context.Context, condition map[string]interface{}, updateAttrs map[string]interface{}) error {
	db := provider.EvaluateDB
	m := database.Account{}
	result := db.Model(m).Where(condition).Updates(updateAttrs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AccountGet(ctx context.Context, condition map[string]interface{}) ([]database.Account, error) {
	db := provider.EvaluateDB
	m := database.Account{}
	account := []database.Account{}
	db = db.Model(m).Select(util.GetJsonFields(m))
	result := db.Where(condition).Where("is_deleted = ?", enums.IsDeletedNo).Find(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return account, nil
}
