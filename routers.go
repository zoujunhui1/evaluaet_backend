package main

import (
	"evaluate_backend/app/handler"
	"github.com/gin-gonic/gin"
)

func routeInit() *gin.Engine {
	r := gin.Default()
	evaluate := r.Group("/evaluate")
	{
		evaluate.GET("/list", handler.GetProductList)
	}
	return r
}
