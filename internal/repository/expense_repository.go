package repository

import "expense-manager/internal/model"

type ExpenseRepository struct {
	expenses []model.Expense
}

func (r *ExpenseRepository) Add(expense model.Expense) model.Expense {
	expense.ID = len(r.expenses) + 1
	r.expenses = append(r.expenses, expense)
	return expense
}

func (r *ExpenseRepository) GetAll() []model.Expense {
	return r.expenses
}

func (r *ExpenseRepository) GetByUserID(userID int) []model.Expense {
	var result []model.Expense
	for _, e := range r.expenses {
		if e.UserID == userID {
			result = append(result, e)
		}
	}
	return result
}

func (r *ExpenseRepository) Delete(id int) bool {
	var newExpenses []model.Expense
	found := false

	for _, e := range r.expenses {
		if e.ID != id {
			newExpenses = append(newExpenses, e)
		} else {
			found = true
		}
	}

	r.expenses = newExpenses
	return found
}

func (r *ExpenseRepository) Total() float64 {
	var total float64
	for _, e := range r.expenses {
		total += float64(e.Amount)
	}
	return total
}

func (r *ExpenseRepository) Update(id int, updated model.Expense) bool {
	for i, e := range r.expenses {
		if e.ID == id {
			r.expenses[i].Title = updated.Title
			r.expenses[i].Amount = updated.Amount
			r.expenses[i].Category = updated.Category
			return true
		}
	}
	return false
}

func (r *ExpenseRepository) GetByID(id int) (model.Expense, bool) {
	for _, e := range r.expenses {
		if e.ID == id {
			return e, true
		}
	}
	return model.Expense{}, false
}

func (r *ExpenseRepository) FilterByUser(userID int, category string, min int) []model.Expense {
	var result []model.Expense

	for _, expense := range r.expenses {
		if expense.UserID == userID {
			if (category == "" || expense.Category == category) &&
				(min == 0 || int(expense.Amount) >= min) {
				result = append(result, expense)
			}
		}
	}

	return result
}
