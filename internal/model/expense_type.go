package model

import "time"

type ExpenseType struct {
	ID           string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title     string    `json:"name" gorm:"uniqueIndex;not null;size:50"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
