package models




type Like struct {
	Id     int `json:"id"`
	PostId int `json:"postId"`
	UserId int `json:"userId"`
}