package main

import (
	"fmt"
	"log"
	"net/http"
	config "videocall/db"
	"videocall/internal/app/product/controller"
	"videocall/internal/app/product/repository"
	"videocall/internal/app/product/routers"
	"videocall/internal/app/product/service"
	"videocall/middleware"
)

func main() {
	db, err := config.ConnectToDb(*config.NewEnvDbConfig())
	if err != nil {
		log.Fatal(err)
	}

	// repository (inject db)
	productRepository := repository.NewProductRepository(db)

	// services (inject db)
	productService := service.NewProductService(productRepository)

	// controller (inject db)
	productController := controller.NewProductController(productService)

	// router (inject controller)
	productRouter := routers.NewProductRouters(productController)

	mux := http.NewServeMux()

	// mendaftarkan router product
	productRouter.RegisterRouter(mux)

	wraperMux := middleware.TraceMiddleware(mux)

	fmt.Println("Server Berjalan Di Port 8000")
	http.ListenAndServe(":8000", wraperMux)
}
