package routes

import (
	"net/http"

	"social-network/controllers"
	"social-network/middleware"
)

func Likes(mux *http.ServeMux) {
	mux.Handle("GET /likes", middleware.Islogged(http.HandlerFunc(controllers.GetLikes)))
	mux.Handle("POST /likes", middleware.Islogged(http.HandlerFunc(controllers.AddLikes)))
	mux.Handle("DELETE /likes", middleware.Islogged(http.HandlerFunc(controllers.DeleteLikes)))

}
