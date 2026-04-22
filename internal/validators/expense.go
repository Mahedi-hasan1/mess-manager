package validators

import (
	"errors"
	"mess-manager/internal/model"

	"github.com/go-playground/validator/v10"
)

func ValidateAddExpense(addExpenseReq *model.Expense) error{
	if addExpenseReq.Type =="" || addExpenseReq.UserID ==""{
		return errors.New("type and user_id is required")	
	}
	validate := validator.New();
	if err := validate.Struct(addExpenseReq); err != nil {
		return err
	}
	return nil
} 