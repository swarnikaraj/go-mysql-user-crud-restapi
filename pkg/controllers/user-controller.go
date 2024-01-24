package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/swarnikaraj/go-mysql-user-crud-restapi/pkg/models"
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}