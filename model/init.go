package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

var DB *sql.DB


func Init() error{
	var err error
	DB,err = sql.Open("mysql",viper.GetString("mysql.source_name"))
	if err != nil {
		return err
	}
	DB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	err = DB.Ping()
	if err != nil {
		return err
	}else{
		log.Println("mysql startup ok")
	}
	return nil
}