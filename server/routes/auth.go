package routes

import (
	"net/http"

	"social-network/controllers"
)

func AuthRoute(mux *http.ServeMux) {
	mux.HandleFunc("POST /signup", controllers.Signup_post)
	mux.HandleFunc("GET /logout", controllers.Logout)
	mux.HandleFunc("POST /login", controllers.Login_post)
}
