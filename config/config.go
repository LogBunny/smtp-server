package config

import "github.com/spf13/viper"

var (
	DB_URI                   = ""
	MIGRATE             bool = false
	WEBHOOK_BACKEND_URL      = ""
)

func LoadCfg() {
	DB_URI = viper.GetString("DB_URI")
	MIGRATE = viper.GetBool("MIGRATE")
	WEBHOOK_BACKEND_URL = viper.GetString("WEBHOOK_BACKEND_URL")
}
