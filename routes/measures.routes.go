package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Dalnec/go-gorm-restapi-psql/db"
	"github.com/Dalnec/go-gorm-restapi-psql/models"
)

func GetMeasuresHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var measures []models.Measure
	db.DB.Find(&measures)
	json.NewEncoder(w).Encode(&measures)
}

func CreateMeasuresHandler(w http.ResponseWriter, r *http.Request) {
	var measures models.Measure
	json.NewDecoder(r.Body).Decode(&measures)
	createdMeasure := db.DB.Create(&measures)
	err := createdMeasure.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&measures)
}