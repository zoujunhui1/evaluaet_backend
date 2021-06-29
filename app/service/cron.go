package service

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"evaluate_backend/app/config"
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/model"
	"evaluate_backend/app/util"

	log "github.com/sirupsen/logrus"
)

func CreateProductQcCodeCron() {
	//1.获取需要生成的数据
	condition := map[string]interface{}{
		"status": enums.ProductStatusQrReady,
	}
	_, productList, err := model.GetProduct(context.Background(), condition, 1, 2)
	if err != nil {
		log.Error("model.GetProduct is error (%+v)", err)
		return
	}
	bindUrl := config.Conf.Custom.BindUrl + "/evaluate/product/info?product_id="
	//2.生成二维码
	for _, v := range productList {
		//2.1:绑定地址
		productID := strconv.FormatInt(v.ProductID, 10)
		bindUrl = bindUrl + productID
		qrCodeUrl, err := CreateQrCodeSrv(bindUrl)
		if err != nil {
			log.Error("CreateQrCodeSrv is error (%+v)", productID)
			continue
		}
		//2.2：二维码图片合并
		qrCodeUrl = strings.Replace(qrCodeUrl, "https", "http", 1)
		urlCode := base64.StdEncoding.EncodeToString([]byte(qrCodeUrl))
		mergeUrl := enums.Template + urlCode + "/gravity/northwest"
		urlParse, _ := url.Parse(qrCodeUrl)
		res, err := http.Get(mergeUrl)
		if err != nil {
			continue
		}
		defer res.Body.Close()
		lastUrl, err := util.ImageUploadCommon(urlParse.Path, res.Body)
		if err != nil {
			log.Error("util.ImageUploadCommon error (%+v)", productID)
			continue
		}
		//3.更新数据库
		if err := model.UpdateProduct(context.Background(), map[string]interface{}{
			"product_id": v.ProductID,
		}, map[string]interface{}{
			"qr_code_url": lastUrl,
			"status":      enums.ProductStatusQrDone,
		}); err != nil {
			log.Error("model.UpdateProduct is error (%+v)", productID)
			continue
		}
	}
	return
}
