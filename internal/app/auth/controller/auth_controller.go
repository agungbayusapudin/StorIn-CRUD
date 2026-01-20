package controller

import (
	"net/http"
	"videocall/internal/app/auth/services"
	"videocall/pkg/response"
)

type AuthController struct {
	service services.AuthServiceInterface
}

func NewAuthController(service services.AuthServiceInterface) *AuthController {
	return &AuthController{service: service}
}

func (ctrl *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), nil)
}

func (ctrl *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), nil)
}

func (ctrl *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), nil)
}
