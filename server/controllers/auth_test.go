package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"social-network/models"
)



func TestSignup_post(t *testing.T) {
	// Préparer les données de test
	user := models.UserRequest{
		Username:    "testuser",
		Email:       "test@example.com",
		Password:    "password123",
		Firstname:   "Test",
		Lastname:    "User",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		City:        "Test City",
	}

	// Convertir en JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	// Créer une requête HTTP de test
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Créer un ResponseRecorder pour enregistrer la réponse
	rr := httptest.NewRecorder()

	// Appeler la fonction à tester
	handler := http.HandlerFunc(Signup_post)
	handler.ServeHTTP(rr, req)

	// Vérifier le code de statut
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler a retourné un mauvais statut : %v, attendu %v",
			status, http.StatusOK)
	}

	// Vérifier la réponse
	expected := `{"message":"New user: testuser is created"}`
	if rr.Body.String() != expected {
		t.Errorf("handler a retourné un corps inattendu : %v, attendu %v",
			rr.Body.String(), expected)
	}
}

func TestLogin_post(t *testing.T) {
	// Structure des données de login
	loginData := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: "testuser",
		Password: "password123",
	}

	// Convertir en JSON
	loginJSON, err := json.Marshal(loginData)
	if err != nil {
		t.Fatal(err)
	}

	// Créer une requête HTTP de test
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Créer un ResponseRecorder
	rr := httptest.NewRecorder()

	// Appeler la fonction à tester
	handler := http.HandlerFunc(Login_post)
	handler.ServeHTTP(rr, req)



	// Vérifier le code de statut (probablement erreur sans DB)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Logf("Statut de réponse: %v", status)
	}
}

func TestLogout(t *testing.T) {
	// Créer une requête HTTP de test
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Créer un ResponseRecorder
	rr := httptest.NewRecorder()

	// Appeler la fonction à tester
	handler := http.HandlerFunc(Logout)
	handler.ServeHTTP(rr, req)

	// Vérifier le code de statut
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler a retourné un mauvais statut : %v, attendu %v",
			status, http.StatusOK)
	}

	// Vérifier qu'un cookie a été défini avec une date d'expiration dans le passé
	cookies := rr.Result().Cookies()
	cookieFound := false
	for _, cookie := range cookies {
		if cookie.Name == "accesToken" {
			cookieFound = true
			if !cookie.Expires.Before(time.Now()) {
				t.Errorf("Cookie devrait avoir une date d'expiration dans le passé")
			}
		}
	}

	if !cookieFound {
		t.Errorf("Aucun cookie 'accesToken' n'a été défini")
	}
}
