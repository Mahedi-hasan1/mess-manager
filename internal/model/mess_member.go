package model

import "time"

type MessMember struct {
	ID string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	MessID string `json:"mess_id" gorm:"index"`
	UserID string `json:"user_id" gorm:"index"`

	User   User `json:"user" gorm:"foreignKey:UserID"`

	Role   int `json:"role"`   // 1/2
	Status int `json:"status"` // 0/1

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
