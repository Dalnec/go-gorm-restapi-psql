package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Dalnec/go-gorm-restapi-psql/db"
	"github.com/Dalnec/go-gorm-restapi-psql/helpers"
	"github.com/Dalnec/go-gorm-restapi-psql/models"
	"github.com/gorilla/mux"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	ID        	uint `json:"id"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Products").Find(&user.Product)

	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err helpers.Error
		err = helpers.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var dbuser models.User
	db.DB.Where("email = ?", user.Email).First(&dbuser)

	//check email is alredy registered or not
	if dbuser.Email != "" {
		var err helpers.Error
		err = helpers.SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = helpers.GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	//insert user details in database
	db.DB.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	var authDetails Authentication
	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err helpers.Error
		err = helpers.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser models.User
	db.DB.Where("email = ?", authDetails.Email).First(&authUser)
	fmt.Println(authUser)
	if authUser.Email == "" {
		var err helpers.Error
		err = helpers.SetError(err, "Correo o Contraseña Incorrectos")
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := helpers.CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err helpers.Error
		err = helpers.SetError(err, "Correo o Contraseña Incorrectos")
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := helpers.GenerateJWT(authUser.Email, authUser.Role)
	if err != nil {
		var err helpers.Error
		err = helpers.SetError(err, "Failed to generate token")
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token Token
	token.ID = authUser.ID
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}