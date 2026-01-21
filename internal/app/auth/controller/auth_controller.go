package controller

import (
	"net/http"

	"videocall/internal/app/auth/schema"
	"videocall/internal/app/auth/services"
	userSchema "videocall/internal/app/users/schema"
	"videocall/pkg/request"
	"videocall/pkg/response"
)

type AuthController struct {
	service services.AuthServiceInterface
}

func NewAuthController(service services.AuthServiceInterface) *AuthController {
	return &AuthController{service: service}
}

func (ctrl *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	var reqUser schema.UserLoginRequest

	err := request.DecodeRequest(r, &reqUser)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	responseData, err := ctrl.service.Login(&reqUser)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), &responseData)
}

func (ctrl *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
		return
	}

	var userReq userSchema.UserRequest
	err := request.DecodeRequest(r, &userReq)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
		return
	}

	err = ctrl.service.Register(&userReq)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
		return
	}

	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), nil)
}

func (ctrl *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	response.WriteResponse(w, r, http.StatusOK, response.CodeSuccess, response.GetMessage(response.CodeSuccess), nil)
}
