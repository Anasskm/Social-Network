package models

import "database/sql"

type Comment struct {
	Id        string `json:"id"`
	Desc      string `json:"desc"`
	Img       string `json:"img"`
	UserId    int    `json:"userId"`
	PostId    int    `json:"postId"`
	CreatedAt string `json:"createdAt"`
}

type CommentQuery struct {
	Id     int            `json:"id"`
	Desc   sql.NullString `json:"desc"`
	Img    sql.NullString `json:"img"`
	UserId int            `json:"userId"`

	CreatedAt  string         `json:"createdAt"`
	PostId     int            `json:"postId"`
	Firstname  string         `json:"firstname"`
	Lastname   string         `json:"lastname"`
	ProfilePic sql.NullString `json:"profilePic"`
}

type CommentResponse struct {
	Id         int    `json:"id"`
	Desc       string `json:"desc"`
	Img        string `json:"img"`
	UserId     int    `json:"userId"`
	CreatedAt  string `json:"createdAt"`
	PostId     int    `json:"postId"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	ProfilePic string `json:"profilePic"`
}

func (CommentQ *CommentQuery) ToCommentResponse() *CommentResponse {
	return &CommentResponse{
		Id:         CommentQ.Id,
		Desc:       CommentQ.Desc.String,
		Img:        CommentQ.Img.String,
		UserId:     CommentQ.UserId,
		CreatedAt:  CommentQ.CreatedAt,
		PostId:     CommentQ.PostId,
		Firstname:  CommentQ.Firstname,
		Lastname:   CommentQ.Lastname,
		ProfilePic: CommentQ.ProfilePic.String,
	}
}
