package main

import (
	"evaluate_backend/app/handler"
	"github.com/gin-gonic/gin"
)

func routeInit() *gin.Engine {
	r := gin.Default()
	evaluate := r.Group("/evaluate/product")
	{
		//产品列表
		evaluate.GET("/list", handler.GetProductList)
		//获取产品详情
		evaluate.GET("/info", handler.GetProductInfo)
		//编辑产品
		evaluate.POST("/edit", handler.EditProduct)
		//删除产品
		evaluate.POST("/del", handler.DelProduct)
	}
	common := r.Group("/common")
	{
		//登陆
		common.POST("/login", handler.Login)
		//退出
		common.POST("/logout", handler.Logout)
		//图片上传
		common.POST("/image_upload", handler.ImageUpload)
	}
	return r
}
