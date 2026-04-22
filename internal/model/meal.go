package model

import "time"

type Meal struct {
	ID           string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID       string    `json:"user_id" gorm:"type:uuid;not null;index"`
	User         User      `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Date    time.Time `json:"date" gorm:"default:CURRENT_TIMESTAMP"`
	Amount     float64   `json:"amount" gorm:"type:decimal(10,6);default:0"`
	TypeID      string    `json:"type_id" gorm:"type_id:uuid;not null;index"`
	Type         MealType `json:"type" gorm:"foreignKey:TypeID;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
