package main

import (
	"fmt"
	"net/http"

	"github.com/PhongVX/golang-rest-api/apis/userapi"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/login", userapi.Login).Methods("POST")
	router.HandleFunc("/api/resource", userapi.GetResource).Methods("POST")

	fmt.Println("Golang Rest API Is Running On Port: 4000")

	err := http.ListenAndServe(":4000", router)

	if err != nil {
		panic(err)
	}

}
