package main

import (
	"net/http"

	"github.com/Dalnec/go-gorm-restapi-psql/db"
	"github.com/Dalnec/go-gorm-restapi-psql/models"
	"github.com/Dalnec/go-gorm-restapi-psql/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// database connection
	db.DBConnection()
	// db.DB.Migrator().DropTable(models.User{})
	db.DB.AutoMigrate(models.Category{})
	db.DB.AutoMigrate(models.Brand{})
	db.DB.AutoMigrate(models.Measure{})
	db.DB.AutoMigrate(models.Product{})
	db.DB.AutoMigrate(models.Prices{})
	db.DB.AutoMigrate(models.User{})

	r:=mux.NewRouter()
	// Index route
	// r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/", routes.HomeHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
    r.Use(mux.CORSMethodMiddleware(r))

	s := r.PathPrefix("/api").Subrouter()

	// // products routes
	s.HandleFunc("/products", routes.GetProductsHandler).Methods("GET")
	s.HandleFunc("/products/{id}", routes.GetProductHandler).Methods("GET")
	s.HandleFunc("/products-associations/{id}", routes.GetProductAssociationHandler).Methods("GET")
	s.HandleFunc("/products", routes.CreateProductHandler).Methods("POST")
	s.HandleFunc("/products-batch", routes.BatchCreateProductsHandler).Methods("POST")
	// s.HandleFunc("/products/{id}/", routes.UpdateProductHandler).Methods("PUT")
	// // Catergories routes
	s.HandleFunc("/categories", routes.GetCategoriesHandler).Methods("GET")
	s.HandleFunc("/categories", routes.CreateCategoriesHandler).Methods("POST")
	// // Brands routes
	s.HandleFunc("/brands", routes.GetBrandsHandler).Methods("GET")
	s.HandleFunc("/brands", routes.CreateBrandsHandler).Methods("POST")
	// // Measures routes
	s.HandleFunc("/measures", routes.GetMeasuresHandler).Methods("GET")
	s.HandleFunc("/measures", routes.CreateMeasuresHandler).Methods("POST")
	// // Measures routes
	// s.HandleFunc("/prices", routes.GetMeasuresHandler).Methods("GET")
	s.HandleFunc("/prices", routes.CreatePricesHandler).Methods("POST")

	// // users routes
	s.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	s.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	s.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	s.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	handler := cors.Default().Handler(r)
	http.ListenAndServe(":4000", handler)
}
