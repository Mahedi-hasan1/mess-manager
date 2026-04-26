package repository

import (
	"mess-manager/internal/model"
	"mess-manager/internal/db"
)

func AddMessMemer(member *model.MessMember) error {
	if err := db.PgDb.Create(member).Error; err != nil {
		return err
	}
	return nil
}

func GetMess(joinCode string) (*model.Mess, error) {
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