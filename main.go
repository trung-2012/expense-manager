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

func updateExpense(w http.ResponseWriter, r *http.Request) {

	var e model.Expense
	err := json.NewDecoder(r.Body).Decode(&e)

	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	ok := manager.UpdateExpense(e.ID, e)

	if !ok {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Expense updated"))
}

func main() {

	http.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			getExpenses(w, r)

		case http.MethodPost:
			createExpense(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}

	})

	http.HandleFunc("/expenses/update", updateExpense)

	http.ListenAndServe(":8080", nil)

}
