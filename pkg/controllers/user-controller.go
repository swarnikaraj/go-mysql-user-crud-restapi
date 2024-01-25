package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/swarnikaraj/go-mysql-user-crud-restapi/pkg/models"
	"github.com/swarnikaraj/go-mysql-user-crud-restapi/pkg/utils"
)

var UserModel models.User
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	 allusers:=UserModel.AllUserGetter()
     res,_:=json.Marshal(allusers)
	 w.Header().Set("Content-Type","application/json")
	 w.WriteHeader(http.StatusOK)
	 w.Write(res)
}



func GetSingleUser(w http.ResponseWriter, r *http.Request) {
   vars:= mux.Vars(r)
   userId:=vars["id"]
   ID,err:=strconv.ParseInt(userId,0,0)
   if err!=nil{
	fmt.Print("Parsing Error Came")
   }
   userfromdb,_:=UserModel.SinglerUserGetter(ID)
   res,_:=json.Marshal(userfromdb)
   w.Header().Set("Content-Type","application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(res)
}



func UserCreator(w http.ResponseWriter, r *http.Request) {
  UserforDb:= &models.User{}
  utils.Parsebody(r,UserforDb)
  u:=UserforDb.UserCreator()
  res,_:=json.Marshal(u)
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusAccepted)
  w.Write(res)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
   vars:=mux.Vars(r)
   userId:=vars["id"]
   ID,err:=strconv.ParseInt(userId,0,0)
   if err!=nil{
	fmt.Print("Error while Parsing")
   }
   removedUser:=UserModel.UserRemover(ID)
   res,_:=json.Marshal(removedUser)
   w.Header().Set("Content-Type","application/json")
   w.WriteHeader(http.StatusOK)
   w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
   updateUser:=&models.User{}
   vars:=mux.Vars(r)
   userId:=vars["id"]
   ID,parseErr:=strconv.ParseInt(userId,0,0)


   if parseErr!=nil{
	fmt.Print("Error while parsing")
   }
   utils.Parsebody(r,updateUser)
   userdetails,db:=UserModel.SinglerUserGetter(ID)

  if updateUser.Name != " "{
	userdetails.Name=updateUser.Name
  }
  if updateUser.Age!= 0 {
    userdetails.Age=updateUser.Age
   }
  if updateUser.Email!=" "{
   userdetails.Email=updateUser.Email
  }
  db.Save(&userdetails)

  res,_:=json.Marshal(userdetails)
  w.Header().Set("Content-Type","application/json")
  w.WriteHeader(http.StatusAccepted)
  w.Write(res)
}

