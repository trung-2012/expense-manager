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

func (s *ExpenseService) DeleteExpense(id int) bool {
	var newExpenses []model.Expense
	found := false

	for _, e := range s.expenses {
		if e.ID != id {
			newExpenses = append(newExpenses, e)
		} else {
			found = true
		}
	}

	s.expenses = newExpenses
	return found
}

func (s *ExpenseService) TotalExpense() float64 {
	var total float64
	for _, e := range s.expenses {
		total += float64(e.Amount)
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

func (s *ExpenseService) GetByID(id int) (model.Expense, bool) {

	for _, e := range s.expenses {
		if e.ID == id {
			return e, true
		}
	}

	return model.Expense{}, false
}

func (s *ExpenseService) FilterExpenses(category string, min int) []model.Expense {
	var result []model.Expense

	for _, e := range s.expenses {
		if category != "" && e.Category != category {
			continue
		}

		if min > 0 && e.Amount < min {
			continue
		}

		result = append(result, e)
	}

	return result
}
