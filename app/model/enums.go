package model

import (
	"context"
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/database"
	"evaluate_backend/app/provider"
	"evaluate_backend/app/util"
)

func GetEnums(ctx context.Context, condition map[string]interface{}, page, pageSize int) (total int64, enumList []database.Enums, err error) {
	db := provider.EvaluateDB
	m := database.Enums{}
	db = db.Model(m).Select(util.GetJsonFields(m)).Where("is_deleted = ?", enums.IsDeletedNo)
	if v, ok := condition["order_by"]; ok {
		db = db.Order(v)
	}
	if v, ok := condition["enum_id"]; ok {
		db = db.Where("enum_id", v)
	}
	if v, ok := condition["father_enum_id"]; ok {
		db = db.Where("father_enum_id", v)
	}
	offset := util.GetOffset(page, pageSize)
	totalQuery := db
	totalQuery.Count(&total)
	result := db.Offset(offset).Limit(pageSize).Order("id desc").Find(&enumList)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	return
}

func GetAllEnums(ctx context.Context, condition map[string]interface{}) (enumList []database.Enums, err error) {
	db := provider.EvaluateDB
	m := database.Enums{}
	db = db.Model(m).Select(util.GetJsonFields(m)).Where("is_deleted = ?", enums.IsDeletedNo)
	if v, ok := condition["order_by"]; ok {
		db = db.Order(v)
	}
	if v, ok := condition["enum_id"]; ok {
		db = db.Where("enum_id", v)
	}
	if v, ok := condition["father_enum_id"]; ok {
		db = db.Where("father_enum_id", v)
	}
	result := db.Order("id desc").Find(&enumList)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}

func AddEnumsModel(ctx context.Context, enumsModel database.Enums) error {
	db := provider.EvaluateDB
	result := db.Create(&enumsModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateEnumsModel(ctx context.Context, condition map[string]interface{}, updateAttrs map[string]interface{}) error {
	db := provider.EvaluateDB
	m := database.Enums{}
	result := db.Model(m).Where(condition).Updates(updateAttrs)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
