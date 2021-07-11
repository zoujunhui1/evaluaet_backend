package service

import (
	"evaluate_backend/app/model"
	"fmt"
)

//生成10000个产品
func CreateProductJobSrv() {
	tmp := []map[string]interface{}{}
	for i := 9901; i <= 10000; i++ {
		insertData := make(map[string]interface{})
		insertData["product_id"] = 100000000 + i
		tmp = append(tmp, insertData)
		if i%100 == 0 {
			_ = model.CreateProduct(nil, tmp)
			tmp = []map[string]interface{}{}
		}
		fmt.Println(i)
	}
}
