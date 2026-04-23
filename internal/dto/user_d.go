package dto

import "time"



type CreateUserRequest struct {
	Username    string `json:"username" validate:"required,min=3,max=50"`
	Email       string `json:"email" validate:"email,max=100"`
	PhoneNumber string `json:"phone_number" gorm:"uniqueIndex;not null;size:100"`
	ImageLink   string `json:"imagelink"`
	Password string `json:"password" validate:"required,min=4"` // Plain password in request
}
type UserResponse struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	Gender         string    `json:"gender"`
	City           string    `json:"city"`
	Email          string    `json:"email"`
	PostCount      int64     `json:"post_count"`
	FollowerCount  int64     `json:"follower_count"`
	FollowingCount int64     `json:"following_count"`
	CreatedAt      time.Time `json:"created_at"`
}
