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
	_, err := cronClient.AddFunc("*/1 * * * * *", service.CreateProductQcCodeCron)
	if err != nil {
		log.Error("cron err(%+v)", err)
	}
	cronClient.Start()

}
