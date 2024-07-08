package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"mymodule/pkg/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}
	
func New(db *gorm.DB) handler {
	return handler{db}
}

func (h handler) AddUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln("Error while sending data", err)
	}

	var user model.User
	json.Unmarshal(body, &user)

	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("created")
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
	userID := vars["id"]

	var user models.
}
