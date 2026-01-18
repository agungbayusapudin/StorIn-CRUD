package main

import (
	"fmt"
	"log"
	"net/http"
	config "videocall/db"
	"videocall/internal/app"

	"videocall/middleware"
)

func main() {
	db, err := config.ConnectToDb(*config.NewEnvDbConfig())
	if err != nil {
		log.Fatal(err)
	}

	mux := app.InitContainer(db)

	wraperMux := middleware.TraceMiddleware(mux)

	fmt.Println("Server Berjalan Di Port 8000")
	http.ListenAndServe(":8000", wraperMux)
}
