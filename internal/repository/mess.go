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
