package transacao

import (
	"fmt"
	"time"

	"github.com/lucioPintanel/personal_finance_project/internal/handler"
)

type createTransacaoRequest struct {
	IDConta      uint      `json:"id_conta"`
	IDTpMoviment uint      `json:"id_tp_mov"`
	IDTpCategory uint      `json:"id_tp_category"`
	Payment      time.Time `json:"paymentAt"`
	Descricao    string    `json:"descricao"`
	Valor        float64   `json:"valor"`
}

func (r *createTransacaoRequest) validate() error {
	handler.Logger.Debugf("Request: %v", r.Descricao)
	if r.Descricao == "" || r.IDConta <= 0 || r.IDTpCategory <= 0 ||
		r.IDTpMoviment <= 0 || r.Valor <= 0.0 {
		return handler.ErrParamIsRequired("descricao", "string")
	}
	return nil
}

type updateTransacaoRequest struct {
	IDConta      uint      `json:"id_conta"`
	IDTpMoviment uint      `json:"id_tp_mov"`
	IDTpCategory uint      `json:"id_tp_category"`
	Payment      time.Time `json:"paymentAt"`
	Descricao    string    `json:"descricao"`
	Valor        float64   `json:"valor"`
}

func (r *updateTransacaoRequest) validate() error {
	handler.Logger.Debugf("Request Update: %v", r.Descricao)
	if r.Descricao != "" || r.IDConta > 0 || r.IDTpCategory > 0 ||
		r.IDTpMoviment > 0 || r.Valor > 0.0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
