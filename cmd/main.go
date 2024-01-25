package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swarnikaraj/go-mysql-user-crud-restapi/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	
    userRouter:=r.PathPrefix("/user").Subrouter()
    routes.UserRoutes(userRouter)

	
	r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	  fmt.Fprintf(w,"Welcome to the User Route")
	})
	r.Handle("/user",r)

    fmt.Print("Server is running on 500")
	err:=http.ListenAndServe(":5000",r)
	if err!=nil{
		panic(err)
	}

}