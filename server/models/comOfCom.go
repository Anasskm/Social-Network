package models

import "database/sql"

type ComOfCom struct {
	Id        string `json:"id"`
	Desc      string `json:"desc"`
	Img       string `json:"img"`
	CommentId int    `json:"commentId"`
	UserId    int    `json:"userId"`
	CreatedAt string `json:"createdAt"`
}

type ComOfComQuery struct {
	Id         int            `json:"id"`
	Desc       sql.NullString `json:"desc"`
	Img        sql.NullString `json:"img"`
	CommentId  int            `json:"commentId"`
	UserId     int            `json:"userId"`
	CreatedAt  string         `json:"createdAt"`
	Firstname  string         `json:"firstname"`
	Lastname   string         `json:"lastname"`
	ProfilePic sql.NullString `json:"profilePic"`
}

type ComOfComResponse struct {
	Id         int    `json:"id"`
	Desc       string `json:"desc"`
	Img        string `json:"img"`
	CommentId  int    `json:"commentId"`
	UserId     int    `json:"userId"`
	CreatedAt  string `json:"createdAt"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	ProfilePic string `json:"profilePic"`
}

func (CommentQ *ComOfComQuery) ToComOfComResponse() *ComOfComResponse {
	return &ComOfComResponse{
		Id:         CommentQ.Id,
		Desc:       CommentQ.Desc.String,
		Img:        CommentQ.Img.String,
		CommentId:  CommentQ.CommentId,
		UserId:     CommentQ.UserId,
		CreatedAt:  CommentQ.CreatedAt,
		Firstname:  CommentQ.Firstname,
		Lastname:   CommentQ.Lastname,
		ProfilePic: CommentQ.ProfilePic.String,
	}
}
