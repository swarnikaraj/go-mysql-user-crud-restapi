package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swarnikaraj/go-mysql-user-crud-restapi/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.UserRoutes(r)
	r.Handle("/",r)
	
    fmt.Print("Server is running on 500")
	err:=http.ListenAndServe(":5000",r)
	if err!=nil{
		panic(err)
	}

}