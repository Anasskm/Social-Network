package models

import "database/sql"

type User struct {
	Id          int            `json:"id"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	Firstname   string         `json:"firstname"`
	Lastname    sql.NullString `json:"lastname"`
	Username    string         `json:"username"`
	DateOfBirth string         `json:"dateOfBirth"`
	CoverPic    sql.NullString `json:"coverPic"`
	ProfilePic  sql.NullString `json:"profilePic"`
	City        sql.NullString `json:"city"`
	CreatedAt   string         `json:"createdAt"`
}

type UserRequest struct {
	Id          int    `json:"id"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	DateOfBirth string `json:"dateOfBirth"`
	CoverPic    string `json:"coverPic"`
	ProfilePic  string `json:"profilePic"`
	City        string `json:"city"`
	CreatedAt   string `json:"createdAt"`
}
type UserResponse struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Username    string `json:"username"`
	DateOfBirth string `json:"dateOfBirth"`
	CoverPic    string `json:"coverPic"`
	ProfilePic  string `json:"profilePic"`
	City        string `json:"city"`
	CreatedAt   string `json:"createdAt"`
}

func (user *User) ToUserResponse() *UserResponse {
	return &UserResponse{
		Id:          user.Id,
		Email:       user.Email,
		Firstname:   user.Firstname,
		Lastname:    user.Lastname.String,
		Username:    user.Username,
		DateOfBirth: user.DateOfBirth,
		CoverPic:    user.CoverPic.String,
		ProfilePic:  user.ProfilePic.String,
		City:        user.City.String,
		CreatedAt:   user.CreatedAt,
	}
}
