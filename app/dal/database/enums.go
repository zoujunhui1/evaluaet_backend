package database

type Enums struct {
	ID           int64  `json:"id"`
	EnumID       int32  `json:"enum_id"`
	EnumName     string `json:"enum_name"`
	FatherEnumID int32  `json:"father_enum_id"`
	Level        int32  `json:"level"`
	IsDeleted    int32  `json:"is_deleted"`
}
