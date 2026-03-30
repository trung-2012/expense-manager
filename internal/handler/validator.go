package handler

import (
	"errors"
	"expense-manager/internal/model"
)

func ValidateExpense(expense model.Expense) error {
	if expense.Title == "" {
		return errors.New("title is required")
	}

	if expense.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if expense.Category == "" {
		return errors.New("category is required")
	}

	return nil
}
