package user

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type createUserResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type deleteUserResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type showUserResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type updateUserResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type listUsersResponse struct {
	handler.BaseResponse
	Data []schemas.UserResponse `json:"data"`
}
