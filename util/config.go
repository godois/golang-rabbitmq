package util

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)


var C *viper.Viper

func init() {
	C = viper.New()
}

func LoadConfig() {
	C.SetConfigType("toml")
	C.SetConfigName("config")    // name of config file (without extension)
	C.AddConfigPath("../config") // call multiple times to add many search paths
	C.AddConfigPath(".")         // optionally look for config in the working directory
	C.AutomaticEnv()             // read in environment variables that match

	if err := C.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		log.Print("[Golang-RabbitMQ - Info] - Configuration file has loaded SUCCESSFULY  ...")
	}
	
}
