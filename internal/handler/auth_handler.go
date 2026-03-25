package handler

import (
	"encoding/json"
	"net/http"

	"expense-manager/internal/auth"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Protected profile success"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GenerateToken("trung", 1)

	if err != nil {
		http.Error(w, "cannot generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
