package model

import "time"

type Mess struct {
	ID   string `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string `json:"name"`

	Members []MessMember `json:"members" gorm:"foreignKey:MessID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
