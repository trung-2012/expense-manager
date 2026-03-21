package service

import "expense-manager/internal/model"

type ExpenseService struct {
	expenses []model.Expense
}

func (s *ExpenseService) AddExpense(expense model.Expense) model.Expense {
	expense.ID = len(s.expenses) + 1
	s.expenses = append(s.expenses, expense)
	return expense
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

func (s *ExpenseService) FilterExpensesByUser(userID int, category string, min int) []model.Expense {
	var result []model.Expense

	for _, expense := range s.expenses {
		if expense.UserID == userID {
			if (category == "" || expense.Category == category) &&
				(min == 0 || int(expense.Amount) >= min) {
				result = append(result, expense)
			}
		}
	}

	return result
}

func (s *ExpenseService) GetByUser(userID int) []model.Expense {
	var result []model.Expense

	for _, expense := range s.expenses {
		if expense.UserID == userID {
			result = append(result, expense)
		}
	}

	return result
}
