package routes

import (
	"net/http"

	"social-network/controllers"
	"social-network/middleware"
)

func ComOfComs(mux *http.ServeMux) {
	mux.Handle("GET /comOfComs", middleware.Islogged(http.HandlerFunc(controllers.GetComOfComs)))

	mux.Handle("POST /comOfComs", middleware.Islogged(http.HandlerFunc(controllers.AddComOfCom)))
}
