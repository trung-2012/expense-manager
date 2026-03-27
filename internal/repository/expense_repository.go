package repository

import (
	"expense-manager/internal/database"
	"expense-manager/internal/model"
)

type ExpenseRepository struct{}

func (r *ExpenseRepository) Add(expense model.Expense) model.Expense {
	result, err := database.DB.Exec(
		"INSERT INTO expenses(title, amount, category, user_id) VALUES (?, ?, ?, ?)",
		expense.Title,
		expense.Amount,
		expense.Category,
		expense.UserID,
	)

	if err != nil {
		return expense
	}

	id, _ := result.LastInsertId()
	expense.ID = int(id)

	return expense
}

func (r *ExpenseRepository) GetAll() []model.Expense {
	rows, err := database.DB.Query(
		"SELECT id, title, amount, category, user_id FROM expenses",
	)

	if err != nil {
		return nil
	}
	defer rows.Close()

	var expenses []model.Expense

	for rows.Next() {
		var expense model.Expense

		rows.Scan(
			&expense.ID,
			&expense.Title,
			&expense.Amount,
			&expense.Category,
			&expense.UserID,
		)

		expenses = append(expenses, expense)
	}

	return expenses
}

func (r *ExpenseRepository) GetByUserID(userID int) []model.Expense {
	rows, err := database.DB.Query(
		"SELECT id, title, amount, category, user_id FROM expenses WHERE user_id=?",
		userID,
	)

	if err != nil {
		return nil
	}
	defer rows.Close()

	var expenses []model.Expense

	for rows.Next() {
		var expense model.Expense

		rows.Scan(
			&expense.ID,
			&expense.Title,
			&expense.Amount,
			&expense.Category,
			&expense.UserID,
		)

		expenses = append(expenses, expense)
	}

	return expenses
}

func (r *ExpenseRepository) Delete(id int) bool {
	result, err := database.DB.Exec(
		"DELETE FROM expenses WHERE id=?",
		id,
	)

	if err != nil {
		return false
	}

	rows, _ := result.RowsAffected()
	return rows > 0
}

func (r *ExpenseRepository) Total() float64 {
	row := database.DB.QueryRow(
		"SELECT SUM(amount) FROM expenses",
	)

	var total float64
	row.Scan(&total)

	return total
}

func (r *ExpenseRepository) Update(id int, updated model.Expense) bool {
	result, err := database.DB.Exec(
		"UPDATE expenses SET title=?, amount=?, category=? WHERE id=?",
		updated.Title,
		updated.Amount,
		updated.Category,
		id,
	)

	if err != nil {
		return false
	}

	rows, _ := result.RowsAffected()
	return rows > 0
}

func (r *ExpenseRepository) GetByID(id int) (model.Expense, bool) {
	row := database.DB.QueryRow(
		"SELECT id, title, amount, category, user_id FROM expenses WHERE id=?",
		id,
	)

	var expense model.Expense

	err := row.Scan(
		&expense.ID,
		&expense.Title,
		&expense.Amount,
		&expense.Category,
		&expense.UserID,
	)

	if err != nil {
		return model.Expense{}, false
	}

	return expense, true
}

func (r *ExpenseRepository) FilterByUser(userID int, category string, min int) []model.Expense {
	query := "SELECT id, title, amount, category, user_id FROM expenses WHERE user_id=?"
	args := []interface{}{userID}

	if category != "" {
		query += " AND category=?"
		args = append(args, category)
	}

	if min > 0 {
		query += " AND amount>=?"
		args = append(args, min)
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var expenses []model.Expense

	for rows.Next() {
		var expense model.Expense

		rows.Scan(
			&expense.ID,
			&expense.Title,
			&expense.Amount,
			&expense.Category,
			&expense.UserID,
		)

		expenses = append(expenses, expense)
	}

	return expenses
}
