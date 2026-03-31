package main

import (
	"net/http"

	"expense-manager/internal/auth"
	"expense-manager/internal/database"
	"expense-manager/internal/handler"
	"expense-manager/internal/middleware"
	"expense-manager/internal/repository"
	"expense-manager/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()

	repo := &repository.ExpenseRepository{}
	svc := service.NewExpenseService(repo)

	handler.ExpenseService = svc

	r := mux.NewRouter()

	r.Use(auth.LoggingMiddleware)
	r.Use(middleware.RecoverMiddleware)

	r.Handle("/expenses", auth.AuthMiddleware(http.HandlerFunc(handler.CreateExpense))).Methods("POST")
	r.Handle("/expenses", auth.AuthMiddleware(http.HandlerFunc(handler.GetExpenses))).Methods("GET")
	r.Handle("/expenses/{id}", auth.AuthMiddleware(http.HandlerFunc(handler.UpdateExpense))).Methods("PUT")
	r.Handle("/expenses/{id}", auth.AuthMiddleware(http.HandlerFunc(handler.DeleteExpense))).Methods("DELETE")
	r.Handle("/expenses/{id}", auth.AuthMiddleware(http.HandlerFunc(handler.GetExpenseByID))).Methods("GET")

	r.HandleFunc("/login", handler.Login).Methods("POST")
	r.Handle("/profile", auth.AuthMiddleware(http.HandlerFunc(handler.Profile))).Methods("GET")
	r.Handle("/logout", auth.AuthMiddleware(http.HandlerFunc(handler.Logout))).Methods("POST")

	http.ListenAndServe(":8080", r)
}
