package service

import (
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/dal/response"
	"evaluate_backend/app/model"
	"evaluate_backend/app/provider"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetProductListSrv(ctx *gin.Context, req *request.GetProductListReq) (*response.GetProductListResp, error) {
	evaluateDB := provider.EvaluateDB
	total, list, err := model.GetProduct(ctx, evaluateDB, req)
	if err != nil {
		return nil, err
	}
	resp := &response.GetProductListResp{}
	resp.Total = total
	resp.PageSize = req.PageSize
	resp.Page = req.Page
	if err := copier.Copy(&resp.List, list); err != nil {
		log.Error("GetProductListSrv copier.Copy is error (%v)", err)
		return nil, err
	}
	return resp, nil

}
