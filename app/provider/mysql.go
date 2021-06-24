package provider

import (
	"evaluate_backend/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var EvaluateDB *gorm.DB

func InitMysql(conf config.Config) (err error) {
	dsn := conf.EvaMysql.Connector()
	EvaluateDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Printf("init mysql err(%+v)", err)
		return err
	}
	return nil
}
