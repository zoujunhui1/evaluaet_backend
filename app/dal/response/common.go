package response

//登陆
type LoginResp struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`  //用户名
	Token string `json:"token"` //token
}

//图片上传
type ImageUploadResp struct {
	Url string `json:"url"`
}

type GetEnumListResp struct {
	Total    int64   `json:"total"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	List     []Enums `json:"enums"`
}

type Enums struct {
	EnumID   int32  `json:"enum_id"`
	EnumName string `json:"enum_name"`
	Level    int32  `json:"level"`
}
