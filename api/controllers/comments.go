package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/models"
)

func GetComments(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	postId := queryParams.Get("postId")

	// db connection and declaring a Slice to hold the results

	db := models.DbConn()
	defer db.Close()

	var comments []models.CommentResponse

	// Execute the query and fetch the results

	q := fmt.Sprintf("SELECT c.*,firstname, lastname, profilePic FROM comments AS c JOIN users as u ON (u.id = c.userId) WHERE c.postId = ? ORDER BY c.createdAt DESC")
	rows, err := db.Query(q, postId)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	// Iterate over the rows and scan the data into the slice

	for rows.Next() {
		var comment models.CommentQuery
		err := rows.Scan(&comment.Id, &comment.Desc, &comment.Img, &comment.UserId, &comment.PostId, &comment.CreatedAt, &comment.Firstname, &comment.Lastname, &comment.ProfilePic)
		if err != nil {
			SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		if comment.Img.Valid {
			comment.Img.String = "http://localhost:3000/uploads/" + comment.Img.String
		}

		comments = append(comments, *comment.ToCommentResponse())
	}

	// Check for errors from iterating over rows.

	if err = rows.Err(); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("*************************************",comment)
	userId, _ := r.Context().Value("userId").(int)
	comment.UserId = userId

	db := models.DbConn()
	defer db.Close()
	query := "INSERT INTO comments (img, `desc`, userId, postId) VALUES (?, ?, ?, ?)"
	if _, err := db.Exec(query, comment.Img, comment.Desc, comment.UserId, comment.PostId); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("A post has been added")
}
