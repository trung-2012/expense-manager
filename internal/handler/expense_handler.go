package handler

import (
	"encoding/json"
	"expense-manager/internal/auth"
	"expense-manager/internal/model"
	"expense-manager/internal/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ExpenseHandler struct {
	service *service.ExpenseService
}

func NewExpenseHandler(service *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

var ExpenseService *service.ExpenseService

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	minStr := r.URL.Query().Get("min")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sort")

	min := 0.0
	if minStr != "" {
		value, _ := strconv.ParseFloat(minStr, 64)
		min = value
	}

	page := 1
	if pageStr != "" {
		value, _ := strconv.Atoi(pageStr)
		page = value
	}

	limit := 5
	if limitStr != "" {
		value, _ := strconv.Atoi(limitStr)
		limit = value
	}

	userID := r.Context().Value(auth.UserIDKey)
	if userID == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	expenses := ExpenseService.GetExpenses(userID.(int), category, min, page, limit, sortBy)

	response := model.APIResponse{
		Message: "success",
		Data:    expenses,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(auth.UserIDKey)
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

	expense = ExpenseService.AddExpense(expense)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.APIResponse{
		Message: "Expense created",
		Data:    expense,
	})
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
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

	ok := ExpenseService.UpdateExpense(id, e)

	if !ok {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode("Expense updated")
}

func GetExpenseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	expense, found := ExpenseService.GetByID(id)

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

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	ok := ExpenseService.DeleteExpense(id)

	if !ok {
		http.Error(w, "Expense not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
