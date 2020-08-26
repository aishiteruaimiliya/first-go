package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)
// 日志init
func LogInfo()  {
	file:="./"+time.Now().Format("2020-01-01")+".log"
	logFile,_:=os.OpenFile(file,os.O_RDWR|os.O_CREATE|os.O_APPEND,0766)
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(logFile)
}

// init 读取配置
func Init() error{
	if err:=Config();err!=nil{
		return err
	}
	LogInfo()
	return nil
}

func Config() error {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
