package user

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type CreateUserResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type DeleteUserResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type ShowUserResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type UpdateUserResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type ListUsersResponse struct {
	handler.BaseResponse
	Data []schemas.UserResponse `json:"data"`
}
