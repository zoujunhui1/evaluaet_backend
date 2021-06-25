package service

import (
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/dal/response"
	"evaluate_backend/app/model"
	"evaluate_backend/app/provider"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"time"
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
	for k, v := range *list {
		createAt := v.CreatedAt.Unix()
		resp.List[k].CreatedAt = time.Unix(createAt, 0).Format("2006-01-02 15:04:05")
	}
	return resp, nil

}
