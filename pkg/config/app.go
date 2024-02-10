package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var db *gorm.DB


func CreateDbConnect(){
	dns:="root:root@tcp(127.0.0.1:3306)/usercrudrest?parseTime=true"
	fmt.Print("Trying connecting db")
	dbcon, err:=gorm.Open(mysql.Open(dns),&gorm.Config{})
    
	if err!=nil{
		panic(err)
	}
	db=dbcon
}


func GetDbConnection() *gorm.DB{
 return db
}


