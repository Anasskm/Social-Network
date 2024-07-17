package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/models"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userId, _ := r.Context().Value("userId").(int)
	// db connection and declaring a Slice to hold the results

	db := models.DbConn()
	defer db.Close()

	var posts []models.PostResponse

	// Execute the query and fetch the results

	q := fmt.Sprintf("SELECT p.*, firstname, lastname, profilePic FROM posts AS p JOIN users as u ON (u.id = p.userId) LEFT JOIN relationships AS r ON (p.userId = r.followedUserId ) WHERE r.followerUserId = ? OR p.userId = ? ORDER BY p.createdAt DESC")
	rows, err := db.Query(q, userId, userId)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	// Iterate over the rows and scan the data into the slice

	for rows.Next() {
		var post models.PostQuery
		err := rows.Scan(&post.Id, &post.Desc, &post.Img, &post.UserId, &post.CreatedAt, &post.Firstname, &post.Lastname, &post.ProfilePic)
		if err != nil {
			SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		if post.Img.Valid {
			post.Img.String = "http://localhost:3000/uploads/" + post.Img.String
		}

		posts = append(posts, *post.ToPostResponse())
	}

	// Check for errors from iterating over rows.

	if err = rows.Err(); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, _ := r.Context().Value("userId").(int)
	post.UserId = userId

	db := models.DbConn()
	defer db.Close()
	query := "INSERT INTO posts (img, `desc`, userId) VALUES (?, ?, ?)"
	if _, err := db.Exec(query, post.Img, post.Desc, post.UserId); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("A post has been added")
}
