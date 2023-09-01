package tipo_item

import (
	"fmt"

	"github.com/lucioPintanel/personal_finance_project/internal/handler"
)

type createTipoItemRequest struct {
	Descricao string `json:"descricao"`
	TipoID    uint   `json:"id_tp"`
}

func (r *createTipoItemRequest) Validate() error {
	handler.Logger.Debugf("Request Descricao: %v", r.Descricao)
	handler.Logger.Debugf("Request TipoRefer: %v", r.TipoID)
	if r.Descricao == "" {
		return handler.ErrParamIsRequired("descricao", "string")
	}
	if r.TipoID <= 0 {
		return handler.ErrParamIsRequired("id_tp", "number")
	}
	return nil
}

type updateTipoItemRequest struct {
	Descricao string `json:"descricao"`
	TipoRefer uint   `json:"id_tp"`
}

func (r *updateTipoItemRequest) Validate() error {
	handler.Logger.Debugf("Request Update: %v", r.Descricao)
	if r.Descricao != "" && r.TipoRefer > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
