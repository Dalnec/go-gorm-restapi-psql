package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Dalnec/go-gorm-restapi-psql/db"
	"github.com/Dalnec/go-gorm-restapi-psql/models"
)

func GetBrandsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var brands []models.Brand
	db.DB.Find(&brands)
	json.NewEncoder(w).Encode(&brands)
}

func CreateBrandsHandler(w http.ResponseWriter, r *http.Request) {
	var brand models.Brand
	json.NewDecoder(r.Body).Decode(&brand)
	createdBrand := db.DB.Create(&brand)
	err := createdBrand.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&brand)
}