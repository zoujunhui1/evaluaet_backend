package enums

const (
	StatusOK             int32 = 0
	ErrorSystemException int32 = 20001
	ErrorInputValidate   int32 = 20002
	ErrorFileUploadFail  int32 = 20003

	ErrorFileUpdateFail int32 = 30001
)

const (
	IsDeletedYes int32  = 1
	IsDeletedNo  int32  = 0
	PwdSalt      string = "Evaluate"
)

//图片水印
const (
	//图片模板
	ImageTemplate string = "http://pingce-new-1256184476.cos.ap-nanjing.myqcloud.com/template_1.png"
	//图片水印
	ImageRemark string = "?watermark/1/image/"
	//位置
	ImageDirection string = "northwest"
)

//文字水印
const (
	//文字水印
	TextRemark string = "?watermark/2/text/"
	//文字水印字体
	FontStyle string = "simhei黑体.ttf"
	//字体大小
	Fontsize string = "40"
	//位置
	Direction string = "south"
	//水平
	Dx string = "30"
	//垂直
	Dy string = "30"
)

//二维码水印

const (
	ProductStatusQrReady         int32 = 10 //待生成二维码
	ProductStatusQrDone          int32 = 11 //二维码生成完成+待编辑
	ProductStatusEditDone        int32 = 20 //编辑完成+文字水印待生成
	ProductStatusTextRemarkDoing int32 = 21 //文字水印生成中
	ProductStatusTextRemarkDone  int32 = 22 //文字水印生成完成
)

var (
	//对提示信息强制指定
	FrontMessage = map[int32]string{
		// common
		StatusOK:             "success",
		ErrorSystemException: "网络异常，请稍后重试",
		ErrorInputValidate:   "请检查输入是否正确",
		ErrorFileUploadFail:  "文件上传失败",
		ErrorFileUpdateFail:  "更新失败",
	}
)

type StatusCode int32

// Message returns the default message for status code
func (status StatusCode) Message() string {
	msg, ok := FrontMessage[int32(status)]
	if !ok {
		return ""
	}
	return msg
}

// Code returns integer code
func (status StatusCode) Code() int32 {
	return int32(status)
}
