package service

import (
	"expense-manager/internal/model"
	"expense-manager/internal/repository"
)

type ExpenseService struct {
	repo repository.ExpenseRepository
}

func (s *ExpenseService) AddExpense(expense model.Expense) model.Expense {
	return s.repo.Add(expense)
}

func (s *ExpenseService) GetAll() []model.Expense {
	return s.repo.GetAll()
}

func (s *ExpenseService) DeleteExpense(id int) bool {
	return s.repo.Delete(id)
}

func (s *ExpenseService) TotalExpense() float64 {
	return s.repo.Total()
}

func (s *ExpenseService) UpdateExpense(id int, updated model.Expense) bool {
	return s.repo.Update(id, updated)
}

func (s *ExpenseService) GetByID(id int) (model.Expense, bool) {
	return s.repo.GetByID(id)
}

func (s *ExpenseService) FilterExpensesByUser(userID int, category string, min int) []model.Expense {
	return s.repo.FilterByUser(userID, category, min)
}

func (s *ExpenseService) GetByUser(userID int) []model.Expense {
	return s.repo.GetByUserID(userID)
}
