package routes

import (
	"net/http"

	"social-network/controllers"
	"social-network/middleware"
)

func LikesCom(mux *http.ServeMux) {
	mux.Handle("GET /comlikes", middleware.Islogged(http.HandlerFunc(controllers.GetLikesCom)))

	// mux.Handle("POST /likes", middleware.Islogged(http.HandlerFunc(controllers.AddComment)))
}
