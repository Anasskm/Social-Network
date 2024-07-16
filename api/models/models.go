package models

type CommentsLike struct {
	Id        int `json:"id"`
	CommentId int `json:"commentId"`
	UserId    int `json:"userId"`
}

type Relationship struct {
	Id             int `json:"id"`
	FollowerUserId int `json:"followerUserId"`
	FollowedUserId int `json:"followedUserId"`
}
