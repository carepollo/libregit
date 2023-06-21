package utils

import (
	"github.com/carepollo/librecode/models"
	"github.com/spf13/viper"
)

// keep env vars in runtime
var GlobalEnv = models.Environment{}

// use dotenv to load .env file declarations into environment variables
func LoadEnv() {
	vp := viper.New()
	vp.SetConfigName("env")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath("../")
	if err := vp.ReadInConfig(); err != nil {
		panic("could not load env vars: " + err.Error())
	}

	if err := vp.Unmarshal(&GlobalEnv); err != nil {
		panic("could not load env vars: " + err.Error())
	}
}
