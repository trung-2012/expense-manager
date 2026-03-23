package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"expense-manager/internal/auth"
	"expense-manager/internal/model"
	"expense-manager/internal/repository"
	"expense-manager/internal/service"

	"github.com/gorilla/mux"
)

var expenseService *service.ExpenseService

func getExpenses(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	minStr := r.URL.Query().Get("min")

	min := 0
	if minStr != "" {
		value, _ := strconv.Atoi(minStr)
		min = value
	}

	userID := r.Context().Value("userID")
	if userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	uid := userID.(int)

	expenses := expenseService.FilterExpensesByUser(uid, category, min)

	response := model.APIResponse{
		Message: "success",
		Data:    expenses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createExpense(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID")
	if userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	expense.UserID = userID.(int)

	if expense.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if expense.Amount <= 0 {
		http.Error(w, "Amount must be greater than 0", http.StatusBadRequest)
		return
	}

	expense = expenseService.AddExpense(expense)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.APIResponse{
		Message: "Expense created",
		Data:    expense,
	})
}

func updateExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, _ := strconv.Atoi(idStr)

	var e model.Expense
	err := json.NewDecoder(r.Body).Decode(&e)

	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if e.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if e.Amount <= 0 {
		http.Error(w, "Amount must be greater than 0", http.StatusBadRequest)
		return
	}

	ok := expenseService.UpdateExpense(id, e)

	if !ok {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode("Expense updated")
}

func getExpenseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	expense, found := expenseService.GetByID(id)

	if !found {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(model.APIResponse{
			Message: "Expense not found",
		})
		return
	}

	json.NewEncoder(w).Encode(model.APIResponse{
		Message: "success",
		Data:    expense,
	})
}

func deleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	ok := expenseService.DeleteExpense(id)

	if !ok {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Protected profile success"))
}

func login(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GenerateToken("trung", 1)

	if err != nil {
		http.Error(w, "cannot generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func main() {
	repo := &repository.ExpenseRepository{}
	expenseService = service.NewExpenseService(repo)

	r := mux.NewRouter()

	r.Use(LoggingMiddleware)

	r.Handle("/expenses", AuthMiddleware(http.HandlerFunc(createExpense))).Methods("POST")
	r.Handle("/expenses", AuthMiddleware(http.HandlerFunc(getExpenses))).Methods("GET")
	r.HandleFunc("/expenses/{id}", updateExpense).Methods("PUT")
	r.HandleFunc("/expenses/{id}", getExpenseByID).Methods("GET")
	r.HandleFunc("/expenses/{id}", deleteExpense).Methods("DELETE")
	r.HandleFunc("/login", login).Methods("POST")
	r.Handle("/profile", AuthMiddleware(http.HandlerFunc(profile))).Methods("GET")

	http.ListenAndServe(":8080", r)
}
