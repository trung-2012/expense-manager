package handler

import (
	"net/http"

	"expense-manager/internal/auth"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Protected profile success",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GenerateToken("trung", 1)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, "cannot generate token")
		return
	}

	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"token":      token,
		"expires_in": "10 minute",
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{
		"message": "logout success",
	})
}
