package model

import (
	"MyfirstGO/pkg/errorno"
	"encoding/json"
	"errors"
	"log"
)

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *User) SelectUserByName(name string) error{
	stmt,err:=DB.Prepare("select user_name,password from user where username=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	row,err:=stmt.Query(name)
	defer row.Close()
	if err != nil {
		return err
	}
	for row.Next(){
		row.Scan(&u.UserName,&u.Password)
	}
	if err:=row.Err();err!=nil{
		return err
	}
	return nil
}

func (u *User) Validate() error {
	if u.UserName==""||u.Password==""{
		return errors.New(errorno.ErrValidation.Message)
	}
	return nil
}

func (u *User) Create() (interface{}, interface{}) {
	id,err:=Insert("insert into user(user_name,password) values(?,?)",u.UserName,u.Password)
	if err != nil {
		return 1,err
	}
	return id,err
}

func (user *User) UserToJSON()string{
	jsonstr,err:=json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	return string(jsonstr)
}

func (user *User)JSONToUser(jsonstr string)error{
	err:=json.Unmarshal([]byte(jsonstr),&user)
	if err != nil {
		return err
	}
	return nil
}