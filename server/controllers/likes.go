package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/models"
)

func GetLikes(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	postId := queryParams.Get("postId")

	// db connection and declaring a Slice to hold the results

	db := models.DbConn()
	defer db.Close()

	likes := []int{}

	// Execute the query and fetch the results

	q := fmt.Sprintf("SELECT userId FROM likes WHERE postId = ?")
	rows, err := db.Query(q, postId)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	// Iterate over the rows and scan the data into the slice

	for rows.Next() {
		var like int
		err := rows.Scan(&like)
		if err != nil {
			SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		likes = append(likes, like)
	}

	// Check for errors from iterating over rows.

	if err = rows.Err(); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	comments := []int{}

	// Execute the query and fetch the results

	q = fmt.Sprintf("SELECT id FROM comments WHERE postId = ?")
	rows, err = db.Query(q, postId)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	// Iterate over the rows and scan the data into the slice

	for rows.Next() {
		var comment int
		err := rows.Scan(&comment)
		if err != nil {
			SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		comments = append(comments, comment)
	}

	// Check for errors from iterating over rows.

	if err = rows.Err(); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	data := []interface{}{likes, comments}

	// Execute the query and fetch the results
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
