package app

import (
	"evaluate_backend/app/config"
	"evaluate_backend/app/provider"
	"fmt"
)

func Init() {
	//配置文件加载
	config.Init()
	//mysql
	if err := provider.InitMysql(config.Conf); err != nil {
		panic(fmt.Sprintf("MySQL Initialization Error: %v", err))
	}
	//cos
	provider.InitCos(config.Conf)
}
