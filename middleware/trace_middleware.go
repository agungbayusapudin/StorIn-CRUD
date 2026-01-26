package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const processId contextKey = "processId"

func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// generate process id
		id := uuid.New().String()

		// memasukan process id ke context
		ctx := context.WithValue(r.Context(), processId, id)

		// Tambahkan logika tracing di sini, misalnya mencatat waktu permintaan
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetId(ctx context.Context) string {
	if id, ok := ctx.Value(processId).(string); ok {
		return id
	}
	return "unknown"
}
