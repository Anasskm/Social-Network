package controllers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"social-network/models"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Data struct {
	Auth  bool
	Title string
}
type Claims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

var (
	JwtKey         = []byte("my_secret_key")
	expirationTime = time.Now().Add(24 * time.Hour)
	data           Data
	emailRegex     = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func SendErrorResponse(w http.ResponseWriter, message string, code int) { // send a response in json format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := ErrorResponse{
		Error: message,
	}
	json.NewEncoder(w).Encode(response)
}

func hashPassword(password string) string { // password hashing
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum([]byte("zebi")))
}

func isEmailValid(email string) bool { // check if email is valid func

	return emailRegex.MatchString(email)
}

func Signup_post(w http.ResponseWriter, r *http.Request) { // signup POST Request handler
	user := models.UserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !isEmailValid(user.Email) {
		SendErrorResponse(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	user.Firstname = strings.Title(strings.ToLower(user.Firstname))
	user.Lastname = strings.Title(strings.ToLower(user.Lastname))
	user.City = strings.Title(strings.ToLower(user.City))

	db := models.DbConn()
	defer db.Close()

	if _, err := db.Exec("INSERT INTO users (username, password, firstname, lastname, email, dateofbirth, city) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.Username, hashPassword(user.Password), user.Firstname, user.Lastname, user.Email, user.DateOfBirth, user.City); err != nil {
		SendErrorResponse(w, "Invalid username or email", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": fmt.Sprintf("New user: %s is created", user.Username),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Login_post(w http.ResponseWriter, r *http.Request) { // login POST Request handler
	user := models.User{}
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := models.DbConn()
	defer db.Close()
	query := "SELECT id, email, password, firstname, lastname, username, dateofbirth, coverpic, profilepic, city, createdat FROM users WHERE email=? OR username=?"
	err := db.QueryRow(query, input.Username, input.Username).Scan(&user.Id, &user.Email, &user.Password, &user.Firstname, &user.Lastname, &user.Username, &user.DateOfBirth, &user.CoverPic, &user.ProfilePic, &user.City, &user.CreatedAt)
	if err != nil {

		if err == sql.ErrNoRows {
			SendErrorResponse(w, "Username Or Email not found", http.StatusBadRequest)
			return
		}
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if hashPassword(input.Password) != user.Password {
		SendErrorResponse(w, "Invalid password", http.StatusBadRequest)
		return
	}
	claims := &Claims{
		Id: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		log.Fatal(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "accesToken",
		Value:    tokenString,
		Expires:  expirationTime,
		Path:     "/",
		HttpOnly: true,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user.ToUserResponse())
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "accesToken",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("logged out")
}
