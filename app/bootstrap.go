package app

import (
	"evaluate_backend/app/service"
	"fmt"
	log "github.com/sirupsen/logrus"

	"evaluate_backend/app/config"
	"evaluate_backend/app/provider"
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
	//cron
	cronClient := provider.InitCron()
	//生成二维码
	//if _, err := cronClient.AddFunc("10 * * * * *",
	//	service.CreateProductQcCodeCron); err != nil {
	//	log.Error("CreateProductQcCodeCron cron err(%+v)", err)
	//}
	//生成文字
	if _, err := cronClient.AddFunc("*/5 * * * * *",
		service.CreateProductTextCron); err != nil {
		log.Error("CreateProductTextCron cron err(%+v)", err)
	}
	cronClient.Start()
}
