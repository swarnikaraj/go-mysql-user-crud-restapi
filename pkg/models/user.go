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


func (u *User) UserCreator() *User{
	
      
		db.Create(&u)
	
	return u
   
}

func (u *User) AllUserGetter() []User{
   var Users []User
   db.Find(&Users)
   return Users
}

func (u *User) SinglerUserGetter(Id int64) (*User, *gorm.DB){
    
	var SingleUser User
	db:=db.Where("ID=?",Id).Find(&SingleUser)
	return &SingleUser,db
}

func (u *User) UserRemover(ID int64) *User{
     var DeletedUser User
	 db.Where("ID=?",ID).Delete(DeletedUser)
	 return &DeletedUser
}





