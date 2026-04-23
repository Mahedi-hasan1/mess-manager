package dto



type AddMessMemerRequest struct {
	MessID string `json:"mess_id" binding:"required,uuid"`
	UserID string `json:"user_id" binding:"required,uuid"`
	Role   int    `json:"role" binding:"required,oneof=1 2"`
}