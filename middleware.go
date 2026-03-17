package main

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method, r.URL.Path, time.Now().Format("15:04:05"))

		next.ServeHTTP(w, r)
	})
}
