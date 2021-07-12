package model

import (
	"context"
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/database"
	"evaluate_backend/app/provider"
	"evaluate_backend/app/util"
	"github.com/pkg/errors"
)

func GetProduct(ctx context.Context, condition map[string]interface{}, page, pageSize int) (total int64, productList []database.Product, err error) {
	db := provider.EvaluateDB
	m := database.Product{}
	db = db.Model(m).Select(util.GetJsonFields(m)).Where("is_deleted = ?", enums.IsDeletedNo)
	if v, ok := condition["order_by"]; ok {
		db = db.Order(v)
	}
	if v, ok := condition["status"]; ok {
		db = db.Where("status", v)
	}
	if v, ok := condition["product_id"]; ok {
		db = db.Where("product_id", v)
	}
	if v, ok := condition["product_id <>"]; ok {
		db = db.Where("product_id <> ?", v)
	}
	if v, ok := condition["product_id >= ?"]; ok {
		db = db.Where("product_id >= ?", v)
	}
	if v, ok := condition["product_id <= ?"]; ok {
		db = db.Where("product_id <= ?", v)
	}
	offset := util.GetOffset(page, pageSize)
	totalQuery := db
	totalQuery.Count(&total)
	result := db.Offset(offset).Limit(pageSize).Find(&productList)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	return
}

func UpdateProduct(ctx context.Context, condition map[string]interface{}, updateAttrs map[string]interface{}) error {
	db := provider.EvaluateDB
	m := database.Product{}
	db = db.Model(m)
	if len(condition) == 0 {
		return errors.Errorf("condition is empty")
	}
	if v, ok := condition["product_id"]; ok {
		db = db.Where("product_id", v)
	}
	if v, ok := condition["product_ids"]; ok {
		db = db.Where("product_id in ?", v)
	}
	result := db.Updates(updateAttrs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateMultiProduct(ctx context.Context, condition map[string]interface{}, updateAttrs map[string]interface{}) error {
	db := provider.EvaluateDB
	m := database.Product{}
	if len(condition) == 0 {
		return errors.Errorf("condition is empty")
	}
	if v, ok := condition["product_ids"]; ok {
		db = db.Where("product_id in (?)", v)
	}
	result := db.Model(m).Updates(updateAttrs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateProduct(ctx context.Context, insertData []map[string]interface{}) error {
	db := provider.EvaluateDB
	m := database.Product{}
	if len(insertData) == 0 {
		return errors.Errorf("insertData is empty")
	}
	result := db.Model(m).Create(insertData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
