package service

import (
	"evaluate_backend/app/const/enums"
	"evaluate_backend/app/dal/request"
	"evaluate_backend/app/dal/response"
	"evaluate_backend/app/model"
	"evaluate_backend/app/util"
	"flag"
	"github.com/golang/freetype"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
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

func CreateImg4Text(text []string, dx, dy int) (string, error) {
	var (
		size    = flag.Float64("size", 12, "font size in points")
		spacing = flag.Float64("spacing", 2, "line spacing (e.g. 2 means double spaced)")
	)
	tt := time.Now().Unix()
	name := "image_text_" + strconv.FormatInt(tt, 10)
	imgFile, _ := os.Create(name)
	defer imgFile.Close()
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	//设置每个点的 RGBA (Red,Green,Blue,Alpha(设置透明度))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			//设置一块 白色(255,255,255)不透明的背景
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	//读取字体数据
	fontBytes, err := ioutil.ReadFile("msyh.ttc")
	if err != nil {
		log.Error("msyh.ttc is error(%+v)", err)
		return "", err
	}
	//载入字体数据
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Error("freetype.ParseFont error(%+v)", err)
		return "", err
	}
	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetFontSize(26)
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色(红色)
	f.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))
	//设置字体的位置
	pt := freetype.Pt(40, 40+int(f.PointToFixed(*size))>>8)
	//写入文字
	for _, v := range text {
		_, err = f.DrawString(v, pt)
		if err != nil {
			log.Error("DrawString error(%+v)", err)
			return "", err
		}
		pt.Y += f.PointToFixed(*size * *spacing)
	}
	//以png 格式写入文件
	err = png.Encode(imgFile, img)
	if err != nil {
		log.Error("png.Encode error(%+v)", err)
		return "", err
	}
	return "", nil
}
