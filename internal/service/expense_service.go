package service

import "expense-manager/internal/model"

type ExpenseService struct {
	expenses []model.Expense
}

func (s *ExpenseService) AddExpense(e model.Expense) {
	s.expenses = append(s.expenses, e)
}

func (s *ExpenseService) GetAll() []model.Expense {
	return s.expenses
}

func (s *ExpenseService) DeleteExpense(id int) {
	var newExpenseService []model.Expense
	for _, e := range s.expenses {
		if id != e.ID {
			newExpenseService = append(newExpenseService, e)
		}
	}
	s.expenses = newExpenseService
}

func (s *ExpenseService) TotalExpense() float64 {
	var total float64
	for _, e := range s.expenses {
		total += e.Amount
	}
	return total
}

func (s *ExpenseService) UpdateExpense(id int, updated model.Expense) bool {
	for i, e := range s.expenses {
		if e.ID == id {
			s.expenses[i].Title = updated.Title
			s.expenses[i].Amount = updated.Amount
			s.expenses[i].Category = updated.Category
			return true
		}
	}
	return false
}
