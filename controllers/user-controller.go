package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/karthik11135/golang-backend/db"
	"github.com/karthik11135/golang-backend/models"
)

type signInBodyModel struct {
	Email string
	Password string
}

func userExists(email string) bool {
	var user models.User
	db.GetDb().Where("email = ?", email).Find(&user)
	fmt.Println("userExists function", user, email)
	return user.Email == email
}

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email, 
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("mysecretpassword"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SignupController(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		json.NewEncoder(w).Encode("Bad request")
		return
	}

	defer r.Body.Close()
	var signupBody models.User
	json.NewDecoder(r.Body).Decode(&signupBody)

	exists := userExists(signupBody.Email)

	if exists {
		json.NewEncoder(w).Encode("Email already exists")
		return;
	}

	db.GetDb().Create(&models.User{Email: signupBody.Email, Password: signupBody.Password, Name: signupBody.Name})

	token, err := createToken(signupBody.Email)

	if err != nil {
		json.NewEncoder(w).Encode("Could not create token")
		return 
	}

	cookie := http.Cookie{Name: "token", Value: token}

	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode("User successfully created")
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var signInBody signInBodyModel

	json.NewDecoder(r.Body).Decode(&signInBody)

	defer r.Body.Close()

	var user = models.User{Email: signInBody.Email}

	if !userExists(user.Email) {
		json.NewEncoder(w).Encode("Wrong password")
		return
	}

	db.GetDb().Find(&user)

	if user.Password != signInBody.Password {
		json.NewEncoder(w).Encode("Wrong password")
		return
	}

	token, _ := createToken(user.Email)

	cookie := http.Cookie{Name: "token", Value: token}

	http.SetCookie(w, &cookie)

	json.NewEncoder(w).Encode("Successfully logged in")
}

func LogoutController (w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "token", Value: "", MaxAge: -1}
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode("User logged out successfully")
}

func AllMyUsers() {
	var users []models.User
	db.GetDb().Find(&users)
}