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
}

type MySqlConfig struct {
	User     string
	Password string
	Host     string
	HostPort string
	Database string
}

func Init() {
	LoadConf()
}

func LoadConf() {
	dir, _ := os.Getwd()
	os.Getenv("")
	path := dir + "/app/conf/config.dev.toml"
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
