package middleware

import (
	"context"
	"fmt"
	"net/http"

	"social-network/controllers"

	"github.com/golang-jwt/jwt/v4"
)

func Islogged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("accesToken")
		if err != nil {
			controllers.SendErrorResponse(w, "Not Logged in!", http.StatusUnauthorized)
			return
		}
		tokenStr := cookie.Value
		claims := &controllers.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return controllers.JwtKey, nil
		})
		if err != nil {
			controllers.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !token.Valid {
			controllers.SendErrorResponse(w, "Token is not valid !", http.StatusForbidden)
			return
		}
		userId := context.WithValue(r.Context(), "userId", claims.Id)

		next.ServeHTTP(w, r.WithContext(userId))
	})
}

func LoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		falsCtx := context.WithValue(r.Context(), "user", false)
		trueCtx := context.WithValue(r.Context(), "user", true)
		cookie, err := r.Cookie("accesToken")
		if err == nil {
			tokenStr := cookie.Value
			claims := &controllers.Claims{}

			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return controllers.JwtKey, nil
			})
			fmt.Println("\n--------------------", claims)
			if err == nil && token.Valid {
				// User is already logged in, redirect to home

				http.Redirect(w, r.WithContext(trueCtx), "/home", http.StatusSeeOther)
				return
			}

		}
		next.ServeHTTP(w, r.WithContext(falsCtx))
	})
}
