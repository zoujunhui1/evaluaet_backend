package util

import (
	"crypto/md5"
	"fmt"
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
