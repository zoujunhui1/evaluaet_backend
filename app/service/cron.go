package service

import (
	"context"
	"encoding/base64"
	"github.com/thoas/go-funk"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"evaluate_backend/app/config"
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/model"
	"evaluate_backend/app/util"

	log "github.com/sirupsen/logrus"
)

func CreateProductQrCodeCron() {
	//1.获取需要生成二维码的数据
	condition := map[string]interface{}{
		"status": enums.ProductStatusQrReady,
	}
	_, productList, err := model.GetProduct(context.Background(), condition, 1, 1)
	if err != nil {
		log.Error("model.GetProduct is error (%+v)", err)
		return
	}
	bindUrl := config.Conf.Custom.BindUrl
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
		urlCode := base64.StdEncoding.EncodeToString([]byte(qrCodeUrl))
		mergeUrl := enums.ImageTemplate + enums.ImageRemark + urlCode +
			"/gravity/" + enums.ImageDirection
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

func CreateProductTextCron() {
	//1.获取需要生成文字水印的数据
	condition := map[string]interface{}{
		"status": enums.ProductStatusEditDone,
	}
	_, productList, err := model.GetProduct(context.Background(), condition, 1, 10)
	if err != nil {
		log.Error("model.GetProduct is error (%+v)", err)
		return
	}
	if len(productList) == 0 {
		return
	}
	proIDs := funk.Get(productList, "ProductID")
	if err := model.UpdateProduct(context.Background(), map[string]interface{}{
		"product_ids": proIDs,
	}, map[string]interface{}{
		"status": enums.ProductStatusTextRemarkDoing,
	}); err != nil {
		log.Error("model.UpdateProduct is error (%+v)", proIDs)
		return
	}
	enumList, err := model.GetAllEnums(context.Background(), map[string]interface{}{})
	enumsMap := make(map[int32]string)
	for _, v := range enumList {
		enumsMap[v.EnumID] = v.EnumName
	}
	for _, v := range productList {
		score := enumsMap[v.Score]                                       //分数
		level := enumsMap[v.Level]                                       //级别
		diameter := strconv.FormatFloat(float64(v.Diameter), 'g', 5, 32) //直径
		thick := strconv.FormatFloat(float64(v.Thick), 'g', 5, 32)       //厚度
		weight := strconv.FormatFloat(float64(v.Weight), 'g', 5, 32)     //重量
		proID := strconv.Itoa(int(v.ProductID))                          //编号id
		text := v.Name + "\n" +
			v.Denomination + " " + v.Material + "\n" +
			v.ProductVersion + "\n" +
			"NPGS " + level + score + "\n" +
			diameter + "*" + thick + "mm " +
			weight + "g\n" + proID //文本
		originUrl := v.QrCodeUrl + enums.TextRemark
		textEncode := base64.URLEncoding.EncodeToString([]byte(text))
		fontStyleEncode := base64.URLEncoding.EncodeToString([]byte(enums.FontStyle))
		mergeUrl := originUrl + textEncode + "/fill/" + fontStyleEncode +
			"/fontsize/" + enums.Fontsize +
			"/dx/" + enums.Dx +
			"/dy/" + enums.Dy +
			"/gravity/" + enums.Direction
		res, err := http.Get(mergeUrl)
		if err != nil {
			continue
		}
		defer res.Body.Close()
		tmpStr := strconv.FormatInt(time.Now().Unix(), 10)
		name := "/text_code/evaluate_text_code_" + tmpStr + ".jpg"
		lastUrl, err := util.ImageUploadCommon(name, res.Body)
		if err != nil {
			continue
		}
		//3.更新数据库
		if err := model.UpdateProduct(context.Background(), map[string]interface{}{
			"product_id": v.ProductID,
		}, map[string]interface{}{
			"text_url": lastUrl,
			"status":   enums.ProductStatusTextRemarkDone,
		}); err != nil {
			log.Error("model.UpdateProduct is error (%+v)", v.ProductID)
			continue
		}
	}

}
