package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"social-network/models"
)

func GetComOfComs(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	commentId := queryParams.Get("commentId")

	// db connection and declaring a Slice to hold the results

	db := models.DbConn()
	defer db.Close()

	var comOfComs []models.ComOfComResponse

	// Execute the query and fetch the results

	q := fmt.Sprintf("SELECT c.*,firstname, lastname, profilePic FROM comOfComs AS c JOIN users as u ON (u.id = c.userId) WHERE c.commentId = ? ORDER BY c.createdAt DESC")
	rows, err := db.Query(q, commentId)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	// Iterate over the rows and scan the data into the slice

	for rows.Next() {
		var comOfCom models.ComOfComQuery
		err := rows.Scan(&comOfCom.Id, &comOfCom.Desc, &comOfCom.Img, &comOfCom.CommentId, &comOfCom.UserId, &comOfCom.CreatedAt, &comOfCom.Firstname, &comOfCom.Lastname, &comOfCom.ProfilePic)
		if err != nil {
			SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		}
		if comOfCom.Img.Valid {
			comOfCom.Img.String = "http://localhost:3000/uploads/" + comOfCom.Img.String
		}
		comOfComs = append(comOfComs, *comOfCom.ToComOfComResponse())
	}

	// Check for errors from iterating over rows.

	if err = rows.Err(); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comOfComs)
}

func AddComOfCom(w http.ResponseWriter, r *http.Request) {
	comment := models.ComOfCom{}
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, _ := r.Context().Value("userId").(int)
	comment.UserId = userId

	db := models.DbConn()
	defer db.Close()
	query := "INSERT INTO comOfComs (img, `desc`, userId, commentId) VALUES (?, ?, ?, ?)"
	if _, err := db.Exec(query, comment.Img, comment.Desc, comment.UserId, comment.CommentId); err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("A post has been added")
}
