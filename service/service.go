package service

import (
	"MyfirstGO/handler"
	"MyfirstGO/model"
	"MyfirstGO/pkg/errorno"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context){
	var r model.User
	if err:=c.Bind(&r);err!=nil{
		handler.SendResponse(c,errorno.ErrBind,nil)
		return
	}
	u:=model.User{
		UserName: r.UserName,
		Password: r.Password,
	}
	if err:=u.Validate();err != nil {
		handler.SendResponse(c,errorno.ErrValidation,nil)
		return
	}
	if _,err := u.Create(); err != nil {
		handler.SendResponse(c, errorno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c,nil,u)
}

func SelectUser(c *gin.Context){
	name:=c.Query("user_name")
	if name==""{
		handler.SendResponse(c,errorno.ErrValidation,nil)
		return
	}
	var user model.User
	if err:=user.SelectUserByName(name);nil!=err{
		fmt.Println(err)
		handler.SendResponse(c,errorno.ErrUserNotFound,nil)
		return
	}
	if err:=user.Validate();err!=nil{
		handler.SendResponse(c,errorno.ErrUserNotFound,nil)
		return
	}
	handler.SendResponse(c,nil,user)
}
