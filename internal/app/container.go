package app

import (
	"database/sql"
	"net/http"

	userController "videocall/internal/app/users/controller"
	userRepo "videocall/internal/app/users/repository"
	userRouters "videocall/internal/app/users/routers"
	userServices "videocall/internal/app/users/services"

	prodController "videocall/internal/app/product/controller"
	prodRepo "videocall/internal/app/product/repository"
	prodRouters "videocall/internal/app/product/routers"
	prodService "videocall/internal/app/product/service"
)

func InitContainer(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// injection masing-masing pada user
	uRepo := userRepo.NewUsersRepository(db)
	uSvc := userServices.NewUserServices(uRepo)
	uCtrl := userController.NewUsersController(uSvc)
	userRouters.NewUserRouters(uCtrl).RegisterUserRouter(mux)

	// injection masing-masing pada product
	pRepo := prodRepo.NewProductRepository(db)
	pSvc := prodService.NewProductService(pRepo)
	pCtrl := prodController.NewProductController(pSvc)
	prodRouters.NewProductRouters(pCtrl).RegisterRouter(mux)

	return mux
}
