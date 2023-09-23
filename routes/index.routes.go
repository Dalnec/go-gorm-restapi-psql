package routes

import (
	"net/http"

	"github.com/Dalnec/go-gorm-restapi-psql/db"
	"github.com/Dalnec/go-gorm-restapi-psql/helpers"
	"github.com/Dalnec/go-gorm-restapi-psql/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}


func InitHandler(w http.ResponseWriter, r *http.Request) {
	
	hash, _ := helpers.GeneratehashPassword("shaddai")
	user := models.User{ FirstName: "Shaddai", LastName: "Shaddai", UserName: "Shaddai", Email:"shaddai@dl.com", Password:hash, Role:"admin"} 
	db.DB.Create(&user)

	brand := models.Brand{ Description: "-", Active: true } 
	db.DB.Create(&brand)

	category := models.Category{ Description: "-", Active: true } 
	db.DB.Create(&category)

	w.Write([]byte("Init Data Filled Successfully!"))
}
