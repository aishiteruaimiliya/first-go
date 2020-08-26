package model

import (
	"fmt"
	"log"
)

// 插入
func Insert(sql string,args ...interface{})(int64,error){
	stmt,err:=DB.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return 1,err
	}
	result,err := stmt.Exec(args...)
	if err != nil {
		return 1,err
	}
	id,err:=result.LastInsertId()
	if err != nil {
		return 1,err
	}
	fmt.Printf("insert ok id is %d",id)
	return id,err
}

// 删除

func Delete(sql string,args...interface{}){
	stmt,err:=DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err,"sql init failed")
	result,err:=stmt.Exec(args...)
	CheckErr(err,"参数添加失败")
	num,err:=result.RowsAffected()
	CheckErr(err,"删除失败")
	fmt.Printf("删除成功，删除了%d行",num)
}

func Update(sql string,args... interface{})  {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err, "SQL语句设置失败")
	result, err := stmt.Exec(args...)
	CheckErr(err, "参数添加失败")
	num, err := result.RowsAffected()
	CheckErr(err,"修改失败")
	fmt.Printf("修改成功，修改行数为%d\n",num)
}


func CheckErr(err error, s string) {
	if err != nil {
		log.Panicln(s)
	}
}
