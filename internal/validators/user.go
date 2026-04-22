package validators

import (
	"errors"
	"mess-manager/internal/dto"
	"mess-manager/internal/util"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateCreateUser(userReq *dto.CreateUserRequest) error {
	// Implement user validation logic
	validate := validator.New()
	if err := validate.Struct(userReq); err != nil {
		customErrorMsg := util.GetCustomValidationMessage(err)
		return errors.New(customErrorMsg)
	}
	if strings.Contains(userReq.Username, "@") {
        return errors.New("username cannot contain '@'")
    }
	return nil
}

