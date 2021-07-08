package response

type GetProductListResp struct {
	Total    int64     `json:"total"`
	PageSize int       `json:"page_size"`
	Page     int       `json:"page"`
	List     []Product `json:"list"`
}

type GetProductInfoResp struct {
	Info *Product `json:"list"`
}

type Product struct {
	ID             int64              `json:"id"`              //自增id
	Name           string             `json:"name"`            //产品名称
	ProductID      int64              `json:"product_id"`      //产品id
	ProductType    string             `json:"product_type"`    //产品类型
	IssueTime      string             `json:"issue_time"`      //发行时间
	Denomination   string             `json:"denomination"`    //面值
	ProductVersion string             `json:"product_version"` //版别
	Weight         float32            `json:"weight"`          //重量
	Thick          float32            `json:"thick"`           //厚度
	Diameter       float32            `json:"diameter"`        //直径
	Score          int32              `json:"score"`           //评级分数
	Level          int32              `json:"level"`           //级别
	IdentifyResult string             `json:"identify_result"` //鉴定结果
	Desc           string             `json:"desc"`            //备注说明
	QrCodeUrl      string             `json:"qr_code_url"`     //二维码地址
	TextUrl        string             `json:"text_url"`        //文本地址
	IsDeleted      int32              `json:"is_deleted"`      //是否删除
	CreatedAt      string             `json:"created_at"`      //创建时间
	Pic            []*ProductAddition `json:"pic"`             //产品图片
}

type ProductAddition struct {
	ProductID  int64  `json:"product_id"`  //产品id
	ProductImg string `json:"product_img"` //产品图片
}
