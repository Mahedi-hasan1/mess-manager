package model

import "time"

type User struct {
	ID           string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Username     string    `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Email        string    `json:"email" gorm:"uniqueIndex;size:100;default:null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	//Role         int8      `json:"role" gorm:"type:smallint;default:1;not null"`
	PhoneNumber  string    `json:"phone_number" gorm:"uniqueIndex;size:100"`
	ImageLink    string    `json:"image_link"`
	MessMembers []MessMember `json:"mess_members" gorm:"foreignKey:UserID"`

	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
