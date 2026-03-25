package main

import (
	"net/http"

	"expense-manager/internal/handler"
	"expense-manager/internal/repository"
	"expense-manager/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	repo := &repository.ExpenseRepository{}
	svc := service.NewExpenseService(repo)

	handler.ExpenseService = svc

	r := mux.NewRouter()

	r.Use(LoggingMiddleware)

	r.Handle("/expenses", AuthMiddleware(http.HandlerFunc(handler.CreateExpense))).Methods("POST")
	r.Handle("/expenses", AuthMiddleware(http.HandlerFunc(handler.GetExpenses))).Methods("GET")
	r.Handle("/expenses/{id}", AuthMiddleware(http.HandlerFunc(handler.UpdateExpense))).Methods("PUT")
	r.Handle("/expenses/{id}", AuthMiddleware(http.HandlerFunc(handler.DeleteExpense))).Methods("DELETE")
	r.HandleFunc("/expenses/{id}", handler.GetExpenseByID).Methods("GET")

	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.Handle("/profile", AuthMiddleware(http.HandlerFunc(handler.Profile))).Methods("GET")

	http.ListenAndServe(":8080", r)
}
