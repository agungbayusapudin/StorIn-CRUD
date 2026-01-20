package routers

import (
	"net/http"
	"videocall/internal/app/auth/controller"
)

type AuthRouter struct {
	controller *controller.AuthController
}

func NewAuthRouters(controller *controller.AuthController) *AuthRouter {
	return &AuthRouter{controller: controller}
}

func (router *AuthRouter) RegisterAuthRouter(mux *http.ServeMux) {
	mux.HandleFunc("/login", router.controller.Login)
	mux.HandleFunc("/register", router.controller.Register)
	mux.HandleFunc("/logout", router.controller.Logout)
}
