package database

import "time"

type Product struct {
	ID             int64     `json:"id"`              //主键id
	Name           string    `json:"name"`            //产品名称
	ProductID      int64     `json:"product_id"`      //产品id
	ProductType    string    `json:"product_type"`    //产品类型
	IssueTime      string    `json:"issue_time"`      //发行时间
	Denomination   string    `json:"denomination"`    //面值
	ProductVersion string    `json:"product_version"` //版别
	Weight         float32   `json:"weight"`          //重量
	Thick          float32   `json:"thick"`           //厚度
	Diameter       float32   `json:"diameter"`        //直径
	Score          int32     `json:"score"`           //评级分数
	Level          int32     `json:"level"`           //级别
	IdentifyResult string    `json:"identify_result"` //鉴定结果
	Desc           string    `json:"desc"`            //备注说明
	IsDeleted      int32     `json:"is_deleted"`      //是否删除
	QrCodeUrl      string    `json:"qr_code_url"`     //二维码地址
	TextUrl        string    `json:"text_url"`        //文本地址
	CreatedAt      time.Time `json:"created_at"`      //创建时间
	UpdatedAt      time.Time `json:"updated_at"`      //更新时间
}
