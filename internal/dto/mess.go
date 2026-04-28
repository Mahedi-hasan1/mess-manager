package dto



type AddMessMemerRequest struct {
	MessID string `json:"mess_id" validate:"required,uuid"`
	UserID string `json:"user_id" validate:"required,uuid"`
	Role     int    `json:"role" validate:"required,oneof=1 2"`
	Status   int    `json:"status" validate:"required,oneof=0 1"`
}

type CreateMessRequest struct{
	Name string `json:"name" binding:"required"`
}