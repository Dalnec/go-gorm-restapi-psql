package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Dalnec/go-gorm-restapi-psql/db"
	"github.com/Dalnec/go-gorm-restapi-psql/models"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var categories []models.Category
	db.DB.Find(&categories)
	json.NewEncoder(w).Encode(&categories)
}

func CreateCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)
	createdCategory := db.DB.Create(&category)
	err := createdCategory.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&category)
}