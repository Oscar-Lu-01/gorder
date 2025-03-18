package main

import (
	"github.com/Oscar-Lu-01/gorder/common/config"
	"github.com/spf13/viper"
	"log"
)

// 编译时候最先执行
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Printf("%v", viper.Get("order"))
}
