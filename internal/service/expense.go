package service

import (
	"mess-manager/internal/model"
	"mess-manager/internal/repository"
)

func AddExpense(reactionAddReq *model.Expense) error {
	return repository.AddExpense(reactionAddReq)
}

