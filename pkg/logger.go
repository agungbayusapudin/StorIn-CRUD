package pkg

import (
	"context"
	"log"
	"videocall/middleware"
)

func info(ctx context.Context, message string) {
	id := middleware.GetId(ctx)
	log.Println("INFO:", message, "Process ID:", id)
}

func errorLog(ctx context.Context, message string) {
	id := middleware.GetId(ctx)
	log.Println("ERROR:", message, "Process ID:", id)
}
