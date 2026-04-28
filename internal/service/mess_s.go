package service

import (
	"mess-manager/internal/db"
	"mess-manager/internal/dto"
	"mess-manager/internal/model"
	"mess-manager/internal/repository"
	"mess-manager/internal/util"

	"gorm.io/gorm"
)

func AddMessMemer(userReq *dto.AddMessMemerRequest) error {

	// member := &model.MessMember{
	// 	MessID: userReq.MessID,
	// 	UserID: userReq.UserID,
	// 	Role:   userReq.Role,
	// }
	// if err := repository.AddMessMemer(,member); err != nil {
	// 	return err
	// }
	return nil
}

func GetJoinCode() string {
	code := util.GenerateJoinCode()
	existing, _ := repository.GetMessByCode(code)
	if existing != nil {
		return GetJoinCode()
	}
	return code
}
func CreateMess(req dto.CreateMessRequest, userID string) (*model.Mess, error) {

	var mess model.Mess

	err := db.PgDb.Transaction(func(tx *gorm.DB) error {

		mess = model.Mess{
			Name:     req.Name,
			JoinCode: util.GenerateJoinCode(),
		}

		if err := repository.CreateMess(tx, &mess); err != nil {
			return err
		}

		member := model.MessMember{
			MessID: mess.ID,
			UserID: userID,
			Role:   2,
			Status: 1,
		}

		if err := repository.AddMessMember(tx, &member); err != nil {
			return nil
		}
		return tx.Preload("Members").
			Preload("Members.User").
			First(&mess, "id = ?", mess.ID).Error
	})

	if err != nil {
		return nil, err
	}

	return &mess, nil
}
