package controllers

import (
	"be-tesis/db"
	"be-tesis/models"
	"be-tesis/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Decode reqbody into struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Fields Validations
	if user.Email == "" {
		utils.RespondError(w, http.StatusBadRequest, "Email is required")
		return
	}

	if user.Name == "" {
		utils.RespondError(w, http.StatusBadRequest, "Name is required")
		return
	}

	if user.DNI == "" {
		utils.RespondError(w, http.StatusBadRequest, "DNI is required")
		return
	}

	if user.Password == "" {
		utils.RespondError(w, http.StatusBadRequest, "Password is required")
		return
	}

	if user.Role == "" {
		utils.RespondError(w, http.StatusBadRequest, "Role is required")
		return
	}

	// Hash password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}
	user.Password = string(hashedPass)

	// Create user in database
	if err := db.DB.Create(&user).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Database error creating user")
		return
	}
	// Send successful response
	utils.Respond(w, http.StatusCreated, user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		DNI      string `json:"dni"`
	}
	json.NewDecoder(r.Body).Decode(&credentials)

	var user models.User
	if err := db.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if user.DNI != credentials.DNI {
		utils.RespondError(w, http.StatusBadRequest, "DNI does not match")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"logout": false,
	})
	tokenString, _ := token.SignedString([]byte(utils.SecretKey))
	json.NewEncoder(w).Encode(tokenString)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(utils.SecretKey), nil
	})
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid token")
		return
	}
	claims := token.Claims.(*jwt.MapClaims)
	(*claims)["logout"] = true
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString([]byte(utils.SecretKey))
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Error signing new token")
		return
	}
	utils.Respond(w, http.StatusOK, map[string]string{
		"token": newTokenString,
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	utils.Respond(w, http.StatusOK, users)
}
