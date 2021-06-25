package enums

const (
	StatusOK             int32 = 0
	ErrorSystemException int32 = 20001
	ErrorInputValidate   int32 = 20002
	ErrorFileUploadFail  int32 = 20003
)

const (
	IsDeletedYes int32 = 1
	IsDeletedNo  int32 = 0
)

var (
	//对提示信息强制指定
	FrontMessage = map[int32]string{
		// common
		StatusOK:             "success",
		ErrorSystemException: "网络异常，请稍后重试",
		ErrorInputValidate:   "请检查输入是否正确",
		ErrorFileUploadFail:  "文件上传失败",
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
