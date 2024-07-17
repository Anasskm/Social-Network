package routes

import (
	"net/http"

	"social-network/controllers"
	"social-network/middleware"
)

func Posts(mux *http.ServeMux) {
	mux.Handle("GET /posts", middleware.Islogged(http.HandlerFunc(controllers.GetPosts)))
	mux.Handle("POST /posts", middleware.Islogged(http.HandlerFunc(controllers.AddPost)))

	mux.Handle("POST /upload", middleware.Islogged(http.HandlerFunc(controllers.Upload)))
}
