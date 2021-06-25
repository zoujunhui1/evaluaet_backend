package main

import (
	"evaluate_backend/app/handler"
	"github.com/gin-gonic/gin"
)

func routeInit() *gin.Engine {
	r := gin.Default()
	evaluate := r.Group("/evaluate")
	{
		//产品列表
		evaluate.GET("/list", handler.GetProductList)
		//编辑产品
		evaluate.POST("/edit", handler.EditProduct)
	}
	return r
}
