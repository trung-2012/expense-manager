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
		WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	expenses := ExpenseService.GetExpenses(userID.(int), category, min, page, limit, sortBy)

	response := model.APIResponse{
		Message: "success",
		Data:    expenses,
	}

	WriteJSON(w, http.StatusOK, response)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense model.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	userID := r.Context().Value(auth.UserIDKey)
	if userID == nil {
		WriteError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	expense.UserID = userID.(int)

	if err := ValidateExpense(expense); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	expense = ExpenseService.AddExpense(expense)

	WriteJSON(w, http.StatusCreated, model.APIResponse{
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
		WriteError(w, http.StatusBadRequest, "Invalid data")
		return
	}

	if err := ValidateExpense(e); err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	ok := ExpenseService.UpdateExpense(id, e)

	if !ok {
		WriteError(w, http.StatusNotFound, "Expense not found")
		return
	}

	WriteJSON(w, http.StatusOK, "Expense updated")
}

func GetExpenseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	expense, found := ExpenseService.GetByID(id)

	if !found {
		WriteJSON(w, http.StatusNotFound, model.APIResponse{
			Message: "Expense not found",
		})
		return
	}

	WriteJSON(w, http.StatusOK, model.APIResponse{
		Message: "success",
		Data:    expense,
	})
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	ok := ExpenseService.DeleteExpense(id)

	if !ok {
		WriteError(w, http.StatusNotFound, "Expense not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
