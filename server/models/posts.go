package models

import "database/sql"

type Post struct {
	Id        string `json:"id"`
	Desc      string `json:"desc"`
	Img       string `json:"img"`
	UserId    int    `json:"userId"`
	CreatedAt string `json:"createdAt"`
}

type PostQuery struct {
	Id         int            `json:"id"`
	Desc       sql.NullString `json:"desc"`
	Img        sql.NullString `json:"img"`
	UserId     int            `json:"userId"`
	CreatedAt  string         `json:"createdAt"`
	Firstname  string         `json:"firstname"`
	Lastname   string         `json:"lastname"`
	ProfilePic sql.NullString `json:"profilePic"`
}

type PostResponse struct {
	Id         int    `json:"id"`
	Desc       string `json:"desc"`
	Img        string `json:"img"`
	UserId     int    `json:"userId"`
	CreatedAt  string `json:"createdAt"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	ProfilePic string `json:"profilePic"`
}

func (postQ *PostQuery) ToPostResponse() *PostResponse {
	return &PostResponse{
		Id:         postQ.Id,
		Desc:       postQ.Desc.String,
		Img:        postQ.Img.String,
		UserId:     postQ.UserId,
		CreatedAt:  postQ.CreatedAt,
		Firstname:  postQ.Firstname,
		Lastname:   postQ.Lastname,
		ProfilePic: postQ.ProfilePic.String,
	}
}
