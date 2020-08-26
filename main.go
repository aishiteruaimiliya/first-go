package main

import (
	"MyfirstGO/config"
	"MyfirstGO/model"
	"MyfirstGO/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err:=config.Init();err!=nil{
		panic(err)
	}
	if err:=model.Init();err!=nil{
		panic(err)
	}
	gin.SetMode(viper.GetString("runmode"))
	g:=gin.New()
	router.InitRouter(g)
	log.Printf("start to listen the incoming requset on http address:%s\n",viper.GetString("addr"))

	if err:=g.Run(viper.GetString("addr"));err!=nil{
		log.Fatal(err)
	}
}
