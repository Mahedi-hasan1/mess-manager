package model

import "time"

type Mess struct {
	ID             string    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name           string    `json:"name" gorm:"type:text"`
	Members        []User    `json:"members" gorm:"many2many:conversation_members;constraint:OnDelete:CASCADE"`
	PendingMembers []User    `json:"pending_members" gorm:"many2many:conversation_pending_members;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

