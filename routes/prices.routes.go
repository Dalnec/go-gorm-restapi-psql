package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Dalnec/go-gorm-restapi-psql/db"
	"github.com/Dalnec/go-gorm-restapi-psql/models"
)

// func GetMeasuresHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
//     w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	var measures []models.Measure
// 	db.DB.Find(&measures)
// 	json.NewEncoder(w).Encode(&measures)
// }

func CreatePricesHandler(w http.ResponseWriter, r *http.Request) {
	var prices models.Prices
	json.NewDecoder(r.Body).Decode(&prices)
	// createdPrices := db.DB.Create(&prices)
	// err := createdPrices.Error

	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(err.Error()))
	// }
	db.DB.Save(&prices)
	json.NewEncoder(w).Encode(&prices)
}