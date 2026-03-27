package service

import (
	"expense-manager/internal/model"
	"expense-manager/internal/repository"
	"sort"
)

type ExpenseService struct {
	repo *repository.ExpenseRepository
}

func NewExpenseService(repo *repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) AddExpense(expense model.Expense) model.Expense {
	return s.repo.Add(expense)
}

func (s *ExpenseService) GetExpenses(userID int, category string, min float64, page int, limit int, sortBy string) []model.Expense {
	expenses := s.repo.FilterByUser(userID, category, int(min))

	var filtered []model.Expense

	for _, e := range expenses {
		if category != "" && e.Category != category {
			continue
		}
		if min > 0 && e.Amount < min {
			continue
		}
		filtered = append(filtered, e)
	}

	if sortBy == "amount" {
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].Amount < filtered[j].Amount
		})
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(filtered) {
		return []model.Expense{}
	}

	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end]
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
