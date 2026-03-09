package main

import (
	"encoding/json"
	"net/http"

	"expense-manager/internal/model"
	"expense-manager/internal/service"
)

var manager = service.ExpenseService{}

func getExpenses(w http.ResponseWriter, r *http.Request) {
	expenses := manager.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

func createExpense(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	manager.AddExpense(expense)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}

func main() {

	http.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			getExpenses(w, r)
		}

		if r.Method == http.MethodPost {
			createExpense(w, r)
		}

	})

	http.ListenAndServe(":8080", nil)

}
