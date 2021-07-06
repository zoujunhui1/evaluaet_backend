package service

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/database"
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/dal/response"
	"evaluate_backend/app/model"
	"evaluate_backend/app/util"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginSrv(ctx *gin.Context, req *request.LoginReq) (*response.LoginResp, error) {
	pwd := util.Md5Sum([]byte(req.Password + enums.PwdSalt))
	//查询用户是否存在
	condition := map[string]interface{}{
		"name":     req.Name,
		"password": pwd,
	}
	account, err := model.AccountGet(ctx, condition)
	if err != nil {
		return nil, err
	}
	if len(account) == 0 {
		return nil, errors.Errorf("user is not exists (%+v)", req)
	}
	user := account[0]
	//生成token
	now := strconv.FormatInt(time.Now().Unix(), 10)
	str := now + enums.PwdSalt
	token := util.Md5Sum([]byte(str))
	if err := model.AccountUpdate(ctx, map[string]interface{}{
		"id": user.ID,
	}, map[string]interface{}{
		"token": token,
	}); err != nil {
		return nil, err
	}
	resp := &response.LoginResp{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}
	return resp, nil
}

func LogoutSrv(ctx *gin.Context, req *request.LogoutReq) error {
	//查询用户是否存在
	condition := map[string]interface{}{
		"id":    req.ID,
		"token": req.Token,
	}
	account, err := model.AccountGet(ctx, condition)
	if err != nil {
		return err
	}
	if len(account) == 0 {
		return errors.Errorf("user is not exists (%+v)", req)
	}
	if err := model.AccountUpdate(ctx, map[string]interface{}{
		"id": req.ID,
	}, map[string]interface{}{
		"token": "",
	}); err != nil {
		return err
	}
	return nil
}

func ImageUploadSrv(ctx *gin.Context, req *request.ImageUploadReq) (*response.ImageUploadResp, error) {
	tmpStr := strconv.FormatInt(time.Now().Unix(), 10)
	name := "/product/evaluate_" + tmpStr + ".jpg"
	fileContent, err := req.Image.Open()
	if err != nil {
		return nil, err
	}
	if fileContent == nil {
		return nil, errors.Errorf("fileContent is nil ")
	}
	defer fileContent.Close()
	url, err := util.ImageUploadCommon(name, fileContent)
	if err != nil {
		return nil, err
	}
	resp := response.ImageUploadResp{Url: url}
	return &resp, nil
}

func CreateQrCodeSrv(bindUrl string) (string, error) {
	//1.生成二维码
	image, err := qrcode.Encode(bindUrl, qrcode.Medium, 128)
	if err != nil {
		return "", err
	}
	//2.上传到cos
	tmpStr := strconv.FormatInt(time.Now().Unix(), 10)
	name := "/qr_code/evaluate_qr_code_" + tmpStr + ".jpg"
	f := strings.NewReader(string(image))
	url, err := util.ImageUploadCommon(name, f)
	if err != nil {
		return "", err
	}
	return url, nil
}

func AddEnumSrv(ctx *gin.Context, req *request.AddEnumReq) error {
	enumMax := 0
	//所传的enum_id是否存在
	_, levelOneData, err := model.GetEnums(ctx, map[string]interface{}{
		"enum_id": req.EnumID,
	}, 1, 1)
	if err != nil {
		return err
	}
	if len(levelOneData) == 0 {
		return errors.Errorf("enum_id is not exists (%+v)", req)
	}
	//获取此enum_id下的枚举最大值
	_, levelTwoData, err := model.GetEnums(ctx, map[string]interface{}{
		"father_enum_id": req.EnumID,
		"order_by":       "enum_id desc",
	}, 1, 1)
	if len(levelTwoData) == 0 {
		enumMax = req.EnumID + 1
	} else {
		enumMax = int(levelTwoData[0].EnumID) + 1
	}
	insertData := database.Enums{
		EnumID:       int32(enumMax),
		EnumName:     req.EnumName,
		FatherEnumID: levelOneData[0].EnumID,
		Level:        2,
	}
	if err := model.AddEnumsModel(ctx, insertData); err != nil {
		return err
	}
	return nil
}

func DelEnumSrv(ctx *gin.Context, req *request.DelEnumReq) error {
	condition := map[string]interface{}{
		"enum_id": req.EnumID,
	}
	updateAttr := map[string]interface{}{
		"is_deleted": enums.IsDeletedYes,
	}
	err := model.UpdateEnumsModel(ctx, condition, updateAttr)
	if err != nil {
		return err
	}
	return nil
}

func GetEnumSrv(ctx *gin.Context, req *request.GetEnumReq) (*response.GetEnumListResp, error) {
	condition := make(map[string]interface{})
	if req.EnumID > 0 {
		condition["enum_id"] = req.EnumID
	}
	total, list, err := model.GetEnums(ctx, condition, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	resp := &response.GetEnumListResp{
		List: []response.Enums{},
	}
	resp.Total = total
	resp.PageSize = req.PageSize
	resp.Page = req.Page
	if len(list) == 0 {
		return resp, nil
	}
	if err := copier.Copy(&resp.List, list); err != nil {
		log.Errorf("GetProductListSrv copier.Copy is error (%+v)", err)
		return nil, err
	}
	return resp, nil
}
