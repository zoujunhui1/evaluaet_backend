package service

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/dal/response"
	"evaluate_backend/app/model"
	"evaluate_backend/app/provider"
	"evaluate_backend/app/util"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginSrv(ctx *gin.Context, req *request.LoginReq) (*response.LoginResp, error) {
	db := provider.EvaluateDB
	pwd := util.Md5Sum([]byte(req.Password + enums.PwdSalt))
	//查询用户是否存在
	condition := map[string]interface{}{
		"name":     req.Name,
		"password": pwd,
	}
	account, err := model.AccountGet(ctx, db, condition)
	if err != nil {
		return nil, err
	}
	if len(account) == 0 {
		return nil, errors.Errorf("user is not exists %(+v)", req)
	}
	user := account[0]
	//生成token
	now := strconv.FormatInt(time.Now().Unix(), 10)
	str := now + enums.PwdSalt
	token := util.Md5Sum([]byte(str))
	if err := model.AccountUpdate(ctx, db, map[string]interface{}{
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
	db := provider.EvaluateDB
	//查询用户是否存在
	condition := map[string]interface{}{
		"id":    req.ID,
		"token": req.Token,
	}
	account, err := model.AccountGet(ctx, db, condition)
	if err != nil {
		return err
	}
	if len(account) == 0 {
		return errors.Errorf("user is not exists %(+v)", req)
	}
	if err := model.AccountUpdate(ctx, db, map[string]interface{}{
		"id": req.ID,
	}, map[string]interface{}{
		"token": "",
	}); err != nil {
		return err
	}
	return nil
}
