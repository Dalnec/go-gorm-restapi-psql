package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dalnec/go-gorm-restapi-psql/db"
	"github.com/Dalnec/go-gorm-restapi-psql/models"
	"github.com/gorilla/mux"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	// db.DB.Preload("Brand").Preload("Category").Find(&products)
	// db.DB.Joins("Brand").Find(&products)
	// db.DB.Find(&products)
	// db.DB.Model(&products).Association("Brand").Find(&products)
	db.DB.Joins("Brand").Joins("Category").Find(&products)
	json.NewEncoder(w).Encode(&products)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	db.DB.First(&product, params["id"])

	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
		return
	}
	json.NewEncoder(w).Encode(&product)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	fmt.Println(product)
	product.Code =strconv.FormatInt(CountCode(), 10)
	createdProduct := db.DB.Create(&product)
	err := createdProduct.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&product)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	db.DB.First(&product, params["id"])

	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
		return
	}

	db.DB.Unscoped().Delete(&product)
	w.WriteHeader(http.StatusOK)
}

func CountCode() int64 {
	var count int64
	var products models.Product

	db.DB.Find(&products).Count(&count)
	return count + 1 
}