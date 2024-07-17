package routes

import (
	"net/http"

	"social-network/controllers"
	"social-network/middleware"
)

func Comments(mux *http.ServeMux) {
	mux.Handle("GET /comments", middleware.Islogged(http.HandlerFunc(controllers.GetComments)))
	mux.Handle("POST /comments", middleware.Islogged(http.HandlerFunc(controllers.AddComment)))
}
