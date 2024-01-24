package models

import (
	"github.com/swarnikaraj/go-mysql-user-crud-restapi/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB
type User struct {
	gorm.Model

	Name string `json:"username"`
	Email string `json:"email"`
	Age int32 `json:"age"`

}

func init(){
	config.CreateDbConnect()
   db=config.GetDbConnection()
   db.AutoMigrate(&User{})

}
