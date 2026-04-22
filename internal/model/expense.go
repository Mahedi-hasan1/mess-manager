package model

import "time"

type Expense struct {
	ID          string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID      string    `json:"user_id" gorm:"type:uuid;not null;index:idx_post_user,unique" validate:"required"`
	User        User      `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Title       string    `json:"title" gorm:"type:text;not null"`
	Amount      float64   `json:"amount" gorm:"type:decimal(10,6);default:0"`
	Date        time.Time `json:"date" gorm:"default:CURRENT_TIMESTAMP"`
	TypeId      string    `json:"type_id" gorm:"type:uuid;not null" validate:"required"`
	ExpenseType User      `json:"-" gorm:"foreignKey:TypeId;references:ID;constraint:OnDelete:CASCADE"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}
