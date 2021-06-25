package util

//获取偏移量
func GetOffset(page, limit int) int {
	if page < 0 {
		return 0
	} else {
		return (page - 1) * limit
	}
}
