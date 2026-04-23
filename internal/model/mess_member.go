package model

import "time"

type MessMember struct {
	ID string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	MessID string `json:"mess_id"`
	Mess   Mess   `json:"mess" gorm:"foreignKey:MessID"`

	UserID string `json:"user_id"`
	User User `json:"user" gorm:"foreignKey:UserID"`

	Role      int       `json:"role"` // 1/2
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
