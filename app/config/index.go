package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

var (
	Conf = Config{}
)

type Config struct {
	EvaMysql *MySqlConfig
	Cos      *CosConfig
	Custom   *Custom
}

type MySqlConfig struct {
	User     string
	Password string
	Host     string
	HostPort string
	Database string
}

type CosConfig struct {
	Host      string
	SecretID  string
	SecretKey string
}

type Custom struct {
	BindUrl string
}

func Init() {
	LoadConf()
}

func LoadConf() {
	dir, _ := os.Getwd()
	path := dir + "/app/conf/config.dev.toml"
	//path := "/Users/zjh/go/src/evaluate_backend/app/conf/config.dev.toml"
	if _, err := toml.DecodeFile(path, &Conf); err != nil {
		fmt.Println(err)
		return
	}
}

func (conf *MySqlConfig) Connector() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.User, conf.Password, conf.Host, conf.HostPort, conf.Database)
	return dsn
}
