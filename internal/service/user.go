package service

import (
	"mess-manager/internal/dto"
	"mess-manager/internal/model"
	"mess-manager/internal/repository"
	"mess-manager/internal/util"
	"strings"
)

func CreateUser(userReq *dto.CreateUserRequest) error {
	hashedPassword, err := util.HashPassword(userReq.Password)
	if err != nil {
		return err
	}
	username := strings.ReplaceAll(userReq.Username, " ", "_")
	user := &model.User{
		Username:     username,
		Email:        userReq.Email,
		PasswordHash: hashedPassword,
		ImageLink: userReq.ImageLink,
	}
	if err := repository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string) (dto.UserResponse, error) {
	user, _ := repository.GetUser("", "", username)
	userRes := dto.UserResponse{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		CreatedAt:      user.CreatedAt,
	}
	return userRes, nil
}

func GetSuggestedUsers(userId string, limit int) ([]model.User, error) {
	return repository.GetSuggestedUsers(userId, limit)
}
