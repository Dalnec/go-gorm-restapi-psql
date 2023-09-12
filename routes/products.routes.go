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

// func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
// 	var products []models.Product
// 	if err := db.DB.Preload("Category").Preload("Brand").Preload("User").Find(&products).Error; err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	// db.DB.Find(&products)
// 	json.NewEncoder(w).Encode(&products)
// }

func GetProductAssociationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	db.DB.First(&product, params["id"])

	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
		return
	}
	// db.DB.Model(&product).Association("Products").Find(&product.ProductID)

	products := []models.Product{}
	db.DB.Preload("Prices").Where("product_id = ?", product.ID).Find(&products)
	product.Products = products
	json.NewEncoder(w).Encode(&product)
}


func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	descriptionFilter := r.URL.Query().Get("description")

	// Build the query condition based on the description filter
	query := db.DB.Preload("Category").Preload("Brand").
		Preload("User").Preload("Prices").Where("product_id IS NULL")

	if descriptionFilter != "" {
		query = query.Where("description LIKE ?", "%"+descriptionFilter+"%")
	}

	if err := query.Order("ID desc").Find(&products).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	// db.DB.Model(&product).Association("Brand").Find(&product.BrandID)
	json.NewEncoder(w).Encode(&product)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	// val, ok := product["ID"]
	if (product.Code == "")  {		
		product.Code = CountCode()
	}

	if product.ProductID == nil {
        product.Product = nil
    }

	if err := db.DB.Save(&product).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    for _, price := range product.Prices {
        if err := db.DB.Save(&price).Error; err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
	// db.DB.Save(&product)
	json.NewEncoder(w).Encode(&product)
}

func BatchCreateProductsHandler(w http.ResponseWriter, r *http.Request) {
	// var product models.Product
	var products []models.Product
	json.NewDecoder(r.Body).Decode(&products)

	for i, product := range products {
		if product.Code == "" {
			// Generate a code if it's not provided
            products[i].Code = CountCode() 
        }
	}
	db.DB.Create(&products)

	json.NewEncoder(w).Encode(&products)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	var product models.Product
	// db.DB.First(&product, params["ID"])
	json.NewDecoder(r.Body).Decode(&product)
	fmt.Println(product)

	if product.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
		return
	}
	db.DB.Save(&product)
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

func CountCode() string {
	var count int64
	var products models.Product

	db.DB.Find(&products).Count(&count)
	var code string = strconv.FormatInt(count + 1 , 10)
	var length = len([]rune(code))
	if length > 4 {
		return code
	} else {
		var cod string
		switch length {
		case 1:
			cod = "000" + code
		case 2:
			cod = "00" + code
		case 3:
			cod = "0" + code
		}
		return cod
	}
}