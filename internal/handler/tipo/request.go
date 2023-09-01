package tipo

import (
	"fmt"

	"github.com/lucioPintanel/personal_finance_project/internal/handler"
)

type createTipoRequest struct {
	Descricao string `json:"descricao"`
}

func (r *createTipoRequest) validate() error {
	handler.Logger.Debugf("Request: %v", r.Descricao)
	if r.Descricao == "" {
		return handler.ErrParamIsRequired("descricao", "string")
	}
	return nil
}

type updateTipoRequest struct {
	Descricao string `json:"descricao"`
}

func (r *updateTipoRequest) validate() error {
	handler.Logger.Debugf("Request Update: %v", r.Descricao)
	if r.Descricao != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
