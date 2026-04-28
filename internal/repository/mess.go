package repository

import (
	"mess-manager/internal/db"
	"mess-manager/internal/model"

	"gorm.io/gorm"
)

func AddMessMember(tx *gorm.DB, member *model.MessMember) error {
	if err := db.PgDb.Create(member).Error; err != nil {
		return err
	}
	return nil
}
func CreateMess(tx *gorm.DB, mess *model.Mess) error {
	if err := db.PgDb.Create(mess).Error; err != nil {
		return err
	}
	return nil
}

func GetMessByCode(joinCode string) (*model.Mess, error) {
	var mess model.Mess
	query := db.PgDb.Model(&model.Mess{})

	if joinCode != "" {
		query = query.Where("join_code = ?", joinCode)
	}
	
	if err := query.First(&mess).Error; err != nil {
		return nil, err
	}
	return &mess, nil
}