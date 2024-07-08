package main

import (
	"log"

	"mymodule/pkg/handlers"
	"mymodule/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	DB := utils.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()
	utils.CountryData()

	router.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)

	log.Println("Api is running")
	http.ListenAndServe(":4000", router)

}
