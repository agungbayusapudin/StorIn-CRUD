package routers

import (
	"net/http"
	"videocall/internal/app/users/controller"
)

type UserRouter struct {
	controller *controller.UserController
}

func NewUserRouters(ctrl *controller.UserController) *UserRouter {
	return &UserRouter{controller: ctrl}
}

func (r *UserRouter) RegisterUserRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", r.controller.GetAllUser)
	mux.HandleFunc("GET /users/{id}", r.controller.GetUserById)
	mux.HandleFunc("POST /users", r.controller.CreateUsers)
	mux.HandleFunc("PUT /users/{id}", r.controller.UpdateUsers)
	mux.HandleFunc("DELETE /users/{id}", r.controller.DeleteUsers)

}
