package service

import (
	"encoding/base64"
	"evaluate_backend/app/util"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/dal/response"
	"evaluate_backend/app/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

func GetProductListSrv(ctx *gin.Context, req *request.GetProductListReq) (*response.GetProductListResp, error) {
	condition := make(map[string]interface{})
	if req.ProductID > 0 {
		condition["product_id"] = req.ProductID
	}
	total, list, err := model.GetProduct(ctx, condition, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	resp := &response.GetProductListResp{
		List: []response.Product{},
	}
	resp.Total = total
	resp.PageSize = req.PageSize
	resp.Page = req.Page
	if err := copier.Copy(&resp.List, list); err != nil {
		log.Errorf("GetProductListSrv copier.Copy is error (%+v)", err)
		return nil, err
	}
	for k, v := range list {
		createAt := v.CreatedAt.Unix()
		resp.List[k].CreatedAt = time.Unix(createAt, 0).Format("2006-01-02 15:04:05")
	}
	return resp, nil
}

func GetProductInfoSrv(ctx *gin.Context, req *request.GetProductInfoReq) (*response.GetProductInfoResp, error) {
	condition := make(map[string]interface{})
	condition["product_id"] = req.ProductID
	_, list, err := model.GetProduct(ctx, condition, 1, 1)
	if err != nil {
		return nil, err
	}
	resp := &response.GetProductInfoResp{
		List: []response.Product{},
	}
	if err := copier.Copy(&resp.List, list); err != nil {
		log.Errorf("GetProductListSrv copier.Copy is error (%v)", err)
		return nil, err
	}
	for k, v := range list {
		createAt := v.CreatedAt.Unix()
		resp.List[k].CreatedAt = time.Unix(createAt, 0).Format("2006-01-02 15:04:05")
	}
	return resp, nil
}

func EditProductSrv(ctx *gin.Context, req *request.EditProductReq) error {
	_, dbData, err := model.GetProduct(ctx, map[string]interface{}{
		"product_id": req.ProductID,
	}, 1, 1)
	if err != nil {
		return err
	}
	if len(dbData) == 0 {
		return errors.Errorf("data is not exists")
	}
	data := dbData[0]
	text := data.Name + "\n" + data.Score + "\n" //文本
	//图片操作
	originUrl := data.QrCodeUrl + enums.TextRemark
	originUrl = strings.Replace(originUrl, "https", "http", 1)
	textEncode := base64.StdEncoding.EncodeToString([]byte(text))
	fontStyleEncode := base64.StdEncoding.EncodeToString([]byte(enums.FontStyle))
	mergeUrl := originUrl + textEncode + "/fill/" + fontStyleEncode +
		"/fontsize/" + enums.Fontsize +
		"/dx/" + enums.Dx +
		"/dy/" + enums.Dy +
		"/gravity/" + enums.Direction
	urlParse, _ := url.Parse(originUrl)
	res, err := http.Get(mergeUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	lastUrl, err := util.ImageUploadCommon(urlParse.Path, res.Body)
	if err != nil {
		log.Error("util.ImageUploadCommon error (%+v)", req.ProductID)

	}
	//生成的产品数量大于0，则说明要更新当前的产品数量（1+req.ProductCount）
	multiProductIDs := []int64{}
	if req.ProductCount > 0 {
		//去数据库里找对应数量的并且未被编辑过的产品
		_, editReadyData, err := model.GetProduct(ctx, map[string]interface{}{
			"status":   enums.ProductStatusEditReady,
			"order_by": "product_id asc",
		}, 1, int(req.ProductCount))
		if err != nil {
			log.Errorf("model.GetProduct error (%v)", err)
			return err
		}
		tmp := funk.Get(editReadyData, "ProductID")
		multiProductIDs = tmp.([]int64)
	}
	multiProductIDs = append(multiProductIDs, req.ProductID)
	condition := map[string]interface{}{
		"product_ids": multiProductIDs,
	}
	updateAttr := make(map[string]interface{})
	updateAttr["status"] = enums.ProductStatusQrReady //更新为已经编辑完成
	updateAttr["qr_code_url"] = lastUrl
	if req.Name != "" {
		updateAttr["name"] = req.Name
	}
	if req.ProductType != "" {
		updateAttr["product_type"] = req.ProductType
	}
	if req.IssueTime != "" {
		updateAttr["issue_time"] = req.IssueTime
	}
	if req.Denomination != "" {
		updateAttr["denomination"] = req.Denomination
	}
	if req.ProductVersion != "" {
		updateAttr["product_version"] = req.ProductVersion
	}
	if req.Weight > 0 {
		updateAttr["weight"] = req.Weight
	}
	if req.Thick > 0 {
		updateAttr["thick"] = req.Thick
	}
	if req.Diameter > 0 {
		updateAttr["diameter"] = req.Diameter
	}
	if req.Score != "" {
		updateAttr["score"] = req.Score
	}
	if req.IdentifyResult != "" {
		updateAttr["identify_result"] = req.IdentifyResult
	}
	if req.Desc != "" {
		updateAttr["desc"] = req.Desc
	}
	err = model.UpdateMultiProduct(ctx, condition, updateAttr)
	if err != nil {
		return err
	}

	return nil
}

func DelProductSrv(ctx *gin.Context, req *request.DelProductReq) error {
	condition := map[string]interface{}{
		"product_id": req.ProductID,
	}
	updateAttr := map[string]interface{}{
		"is_deleted": enums.IsDeletedYes,
	}
	err := model.UpdateProduct(ctx, condition, updateAttr)
	if err != nil {
		return err
	}
	return nil
}
