package service

import (
	"evaluate_backend/app/model"
	"fmt"
)

//生成10000个产品
func CreateProductJobSrv() {
	tmp := []map[string]interface{}{}
	for i := 1; i <= 200; i++ {
		insertData := make(map[string]interface{})
		insertData["product_id"] = 200000000 + i
		tmp = append(tmp, insertData)
		if i%10 == 0 {
			_ = model.CreateProduct(nil, tmp)
			tmp = []map[string]interface{}{}
		}
		fmt.Println(i)
	}
}
