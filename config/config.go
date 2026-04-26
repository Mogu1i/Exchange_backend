// 项目配置信息包
package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenCons  int
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
}

var Appconfig *Config

func InitConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		viper.AddConfigPath("./config")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config:%v", err)
	}

	Appconfig = &Config{}

	if err := viper.Unmarshal(Appconfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	InitDB()
	InitRedis()
}

// InfoTest := Test{
// 	message : "pong",
// }
