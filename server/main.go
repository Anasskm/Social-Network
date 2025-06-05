package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"social-network/middleware"
	"social-network/routes"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	InfoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// CORS middleware setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // allow my front-end app to send requests
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// wrap mux with middleware
	handler := middleware.Logging(mux)
	handler = c.Handler(handler)

	// routes
	routes.AuthRoute(mux)
	routes.Posts(mux)
	routes.Comments(mux)
	routes.ComOfComs(mux)
	routes.Likes(mux)
	routes.LikesCom(mux)

	// http.Server Setup
	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
		IdleTimeout:    120 * time.Second,
	}

	InfoLog.Println("Server started on: http://localhost" + server.Addr)
	log.Fatal(server.ListenAndServe())
}
