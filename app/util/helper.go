package util

import (
	"context"
	"crypto/md5"
	"evaluate_backend/app/config"
	"evaluate_backend/app/provider"
	"fmt"
	"io"
)

//获取偏移量
func GetOffset(page, limit int) int {
	if page < 0 {
		return 0
	} else {
		return (page - 1) * limit
	}
}

//md5
func Md5Sum(Byte []byte) string {
	has := md5.Sum(Byte)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func ImageUploadCommon(name string, content io.Reader) (string, error) {
	_, err := provider.CosClient.Object.Put(context.Background(), name, content, nil)
	if err != nil {
		return "", err
	}
	imgUrl := config.Conf.Cos.Host + "/" + name
	return imgUrl, nil
}
