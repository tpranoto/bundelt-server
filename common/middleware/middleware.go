package middleware

import (
	"log"
	"net/http"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

func MultiMiddlwares(h http.HandlerFunc, fn ...middleware) http.HandlerFunc {
	if len(fn) < 1 {
		return h
	}

	handler := h

	for i := len(fn) - 1; i > 0; i-- {
		handler = fn[i](handler)
	}

	return handler
}

func AccessLog(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
