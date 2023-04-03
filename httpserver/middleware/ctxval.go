package middleware

import (
	"context"
	"net/http"
)

func WithValue[T any](key any, val T) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx2 := context.WithValue(ctx, key, val)
			r2 := r.WithContext(ctx2)
			next.ServeHTTP(w, r2)
		})
	}
}
