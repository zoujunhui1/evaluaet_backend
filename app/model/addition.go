package model

import (
	"context"
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/database"
	"evaluate_backend/app/provider"
	"evaluate_backend/app/util"
)

func GetProductAddition(ctx context.Context, condition map[string]interface{}) (productList []database.ProductAddition, err error) {
	db := provider.EvaluateDB
	m := database.ProductAddition{}
	db = db.Model(m).Select(util.GetJsonFields(m)).Where("is_deleted = ?", enums.IsDeletedNo)
	if v, ok := condition["order_by"]; ok {
		db = db.Order(v)
	}
	if v, ok := condition["product_id"]; ok {
		db = db.Where("product_id", v)
	}
	if v, ok := condition["product_id <>"]; ok {
		db = db.Where("product_id <> ?", v)
	}
	result := db.Order("id asc").Find(&productList)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
