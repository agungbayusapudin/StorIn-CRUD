package app

import (
	"database/sql"
	"net/http"

	authController "videocall/internal/app/auth/controller"
	authRouters "videocall/internal/app/auth/routers"
	authServices "videocall/internal/app/auth/services"

	userController "videocall/internal/app/users/controller"
	userRepo "videocall/internal/app/users/repository"
	userRouters "videocall/internal/app/users/routers"
	userServices "videocall/internal/app/users/services"

	prodController "videocall/internal/app/product/controller"
	prodRepo "videocall/internal/app/product/repository"
	prodRouters "videocall/internal/app/product/routers"
	prodService "videocall/internal/app/product/service"

	orderController "videocall/internal/app/orders/controller"
	orderRepo "videocall/internal/app/orders/repository"
	orderRouters "videocall/internal/app/orders/routers"
	orderService "videocall/internal/app/orders/services"

	billingController "videocall/internal/app/billing/controller"
	billingRepo "videocall/internal/app/billing/repository"
	billingRouters "videocall/internal/app/billing/routers"
	billingService "videocall/internal/app/billing/services"
)

func InitContainer(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// injection masing-masing pada user
	uRepo := userRepo.NewUsersRepository(db)
	uSvc := userServices.NewUserServices(uRepo)
	uCtrl := userController.NewUsersController(uSvc)
	userRouters.NewUserRouters(uCtrl).RegisterUserRouter(mux)

	// injection masing-masing pada auth
	authSvc := authServices.NewAuthService(uRepo, "")
	authCtrl := authController.NewAuthController(authSvc)
	authRouters.NewAuthRouters(authCtrl).RegisterAuthRouter(mux)

	// injection masing-masing pada product
	pRepo := prodRepo.NewProductRepository(db)
	pSvc := prodService.NewProductService(pRepo)
	pCtrl := prodController.NewProductController(pSvc)
	prodRouters.NewProductRouters(pCtrl).RegisterRouter(mux)

	// injection masing-masing pada order
	oRepo := orderRepo.NewOrderRepository(db)
	oRepoDetail := orderRepo.NewOrderDetailRepository(db)

	oSvc := orderService.NewOrderService(oRepo, oRepoDetail)
	oCtrl := orderController.NewOrderControllerInterface(oSvc)
	orderRouters.NewOrderRouters(oCtrl).RegisterOrderRouter(mux)

	// injection masing-masing pada billing
	billRepo := billingRepo.NewBillingRepository(db)
	billSvc := billingService.NewBillingService(billRepo)
	billCtrl := billingController.NewBillingController(billSvc)
	billingRouters.NewBillingRouter(billCtrl).RegisterBillingRouter(mux)

	return mux
}
