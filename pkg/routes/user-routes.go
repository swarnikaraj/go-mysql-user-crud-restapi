package routes

import (
	"github.com/gorilla/mux"
	"github.com/swarnikaraj/go-mysql-user-crud-restapi/pkg/controllers"
)

var UserRoutes = func(router *mux.Router){
    
	router.HandleFunc("/getAllUsers",controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/getSingleUser", controllers.GetSingleUser).Methods("GET")
	router.HandleFunc("/createUser", controllers.UserCreator).Methods("POST")
	router.HandleFunc("/updateUser/{id}",controllers.UpdateUser).Methods("PATCH")
	router.HandleFunc("/deleteUser/{id}", controllers.DeleteUser).Methods("DELETE")

}

