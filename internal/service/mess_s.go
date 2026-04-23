package service

import (
	"mess-manager/internal/dto"
	"mess-manager/internal/model"
	"mess-manager/internal/repository"
)

func AddMessMemer(userReq *dto.AddMessMemerRequest) error {

	member := &model.MessMember{
		MessID: userReq.MessID,
		UserID: userReq.UserID,
		Role:   userReq.Role,
	}
	if err := repository.AddMessMemer(member); err != nil {
		return err
	}
	return nil
}
