package model

import "time"

type Mess struct {
	ID   string `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	JoinCode string `json:"join_code" gorm:"type:varchar(6);not null;uniqueIndex" binding:"required"`
	Name string `json:"name" gorm:"not null" binding:"required"`
	Members []MessMember `json:"members" gorm:"foreignKey:MessID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
