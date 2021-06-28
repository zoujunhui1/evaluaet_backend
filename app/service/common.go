package service

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/dal/response"
	"evaluate_backend/app/model"
	"evaluate_backend/app/util"
	"github.com/pkg/errors"
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
		ID:   user.ID,
		Name: user.Name,
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
	name := "/product/evaluate_" + tmpStr + ".png"
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
	png, err := qrcode.Encode(bindUrl, qrcode.Medium, 128)
	if err != nil {
		return "", err
	}
	//2.上传到cos
	tmpStr := strconv.FormatInt(time.Now().Unix(), 10)
	name := "qr_code/evaluate_qr_code_" + tmpStr + ".png"
	f := strings.NewReader(string(png))
	url, err := util.ImageUploadCommon(name, f)
	if err != nil {
		return "", err
	}
	return url, nil
}
