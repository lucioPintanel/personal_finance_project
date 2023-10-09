package bank_account

import (
	"fmt"

	"github.com/lucioPintanel/personal_finance_project/internal/handler"
)

type createBankAccountRequest struct {
	IDUser    uint    `json:"id_user"`
	Descricao string  `json:"descricao"`
	NroConta  string  `json:"nro_conta"`
	Balance   float64 `json:"balance"`
}

func (r *createBankAccountRequest) validate() error {
	handler.Logger.Debugf("Request Update: %v", r.NroConta)
	if r.Descricao != "" && r.NroConta != "" && r.IDUser > 0 && r.Balance > 0.0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

type updateBankAccountRequest struct {
	IDUser    uint    `json:"id_user"`
	Descricao string  `json:"descricao"`
	NroConta  string  `json:"nro_conta"`
	Balance   float64 `json:"balance"`
}

func (r *updateBankAccountRequest) validate() error {
	handler.Logger.Debugf("Request Update: %v", r.NroConta)
	if r.Descricao != "" && r.NroConta != "" && r.IDUser > 0 && r.Balance > 0.0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
