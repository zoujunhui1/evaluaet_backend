package main

import (
	"evaluate_backend/app/handler"
	"evaluate_backend/app/middleware"
	"github.com/gin-gonic/gin"
)

func routeInit() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	evaluate := r.Group("/evaluate/product")
	{
		//产品列表
		evaluate.GET("/list", handler.GetProductList)
		//范围查询
		evaluate.GET("/range_list", handler.GetProductRangeList)
		//获取产品详情
		evaluate.GET("/info", handler.GetProductInfo)
		//编辑产品
		evaluate.POST("/edit", handler.EditProduct)
		//删除产品
		evaluate.POST("/del", handler.DelProduct)
		//产品图片下载
		evaluate.GET("/image_download", handler.ImageDownload)
	}
	common := r.Group("/common")
	{
		//登陆
		common.POST("/login", handler.Login)
		//退出
		common.POST("/logout", handler.Logout)
		//图片上传
		common.POST("/image_upload", handler.ImageUpload)
		//添加枚举
		common.POST("/enum/add", handler.AddEnum)
		//删除枚举
		common.POST("/enum/del", handler.DelEnum)
		//查询枚举
		common.GET("/enum/list", handler.GetEnumList)
	}
	return r
}
