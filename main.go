package main

import (
	"expense-manager/internal/model"
	"expense-manager/internal/service"
	"fmt"
)

func main() {

	manager := service.ExpenseService{}

	manager.AddExpense(model.Expense{1, "Lunch", 50, "Food"})
	manager.AddExpense(model.Expense{2, "Taxi", 30, "Transport"})
	manager.AddExpense(model.Expense{3, "Coffee", 20, "Drink"})

	fmt.Println("Total:", manager.TotalExpense())

	manager.DeleteExpense(2)

	fmt.Println("Total after delete:", manager.TotalExpense())
}
