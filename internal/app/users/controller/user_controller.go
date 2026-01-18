package controller

import (
	"net/http"
	"strconv"
	"videocall/internal/app/users/schema"
	"videocall/internal/app/users/services"
	"videocall/pkg/request"
	"videocall/pkg/response"
)

type UserController struct {
	service services.UserServiceInterface
}

func NewUsersController(service services.UserServiceInterface) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) CreateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	var reqUser schema.Users

	err := request.DecodeRequest(r, &reqUser)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	err = ctrl.service.CreateUsers(&reqUser)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), nil)
}

func (ctrl *UserController) UpdateUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	var reqUser schema.UserRequest

	err = request.DecodeRequest(r, &reqUser)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	err = ctrl.service.UpdateUsers(id, &reqUser)
	if err != nil {
		response.WriteResponse(w, r, http.StatusInternalServerError, response.CodeInternalError, response.GetMessage(response.CodeInternalError), nil)
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), nil)

}

func (ctrl *UserController) DeleteUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	err = ctrl.service.DeleteUsers(id)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), nil)
}

func (ctrl *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	userData, err := ctrl.service.GetUserById(id)
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), &userData)
}

func (ctrl *UserController) GetAllUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		response.WriteResponse(w, r, http.StatusMethodNotAllowed, response.CodeMethodNotAllowed, response.GetMessage(response.CodeMethodNotAllowed), nil)
	}

	userData, err := ctrl.service.GetAllUser()
	if err != nil {
		response.WriteResponse(w, r, http.StatusBadRequest, response.CodeBadRequest, response.GetMessage(response.CodeBadRequest), nil)
	}

	response.WriteResponse(w, r, http.StatusCreated, response.CodeSuccess, response.GetMessage(response.CodeSuccess), userData)

}
