package validators

import (
	"errors"
	"mess-manager/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateAddMessMember(req *dto.AddMessMemerRequest) error {
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		return err;
	}
	if req.MessID == req.UserID {
		return errors.New("mess_id and user_id cannot be same")
	}

	return nil
}
func ValidateCreateMess(req *dto.CreateMessRequest) error {
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		return err;
	}

	return nil
}