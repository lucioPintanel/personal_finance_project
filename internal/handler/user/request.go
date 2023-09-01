package user

import (
	"fmt"

	"github.com/lucioPintanel/personal_finance_project/internal/handler"
)

type createUserRequest struct {
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
	Cpf   string `json:"cpf"`
}

func (r *createUserRequest) Validate() error {
	handler.Logger.Debugf("Request: %v", r.Nome)
	if r.Nome == "" && r.Senha == "" && r.Cpf == "" {
		return handler.ErrParamIsRequired("nome", "string")
	}
	return nil
}

type updateUserRequest struct {
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}

func (r *updateUserRequest) Validate() error {
	handler.Logger.Debugf("Request Update: %v", r.Nome)
	if r.Nome != "" && r.Senha != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
