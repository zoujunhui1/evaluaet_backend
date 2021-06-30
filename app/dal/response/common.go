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
