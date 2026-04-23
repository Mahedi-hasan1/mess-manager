package dto

import(
	"time"
)

type AddMealRequest struct {
	TypeID string    `json:"type_id" binding:"required,uuid"`
	Date   time.Time `json:"date" binding:"required"`
	Amount float64   `json:"amount" binding:"required,gte=0"`
}