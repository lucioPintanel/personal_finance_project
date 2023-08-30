package handler

import "fmt"

type CreateTipoRequest struct {
	Descricao string `json:"descricao"`
}

type CreateTipoItemRequest struct {
	Descricao string `json:"descricao"`
	TipoID    uint   `json:"id_tp"`
}

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateTipoRequest) Validate() error {
	Logger.Debugf("Request: %v", r.Descricao)
	if r.Descricao == "" {
		return ErrParamIsRequired("descricao", "string")
	}
	return nil
}

func (r *CreateTipoItemRequest) Validate() error {
	Logger.Debugf("Request Descricao: %v", r.Descricao)
	Logger.Debugf("Request TipoRefer: %v", r.TipoID)
	if r.Descricao == "" {
		return ErrParamIsRequired("descricao", "string")
	}
	if r.TipoID <= 0 {
		return ErrParamIsRequired("id_tp", "number")
	}
	return nil
}

type UpdateTipoRequest struct {
	Descricao string `json:"descricao"`
}

func (r *UpdateTipoRequest) Validate() error {
	Logger.Debugf("Request Update: %v", r.Descricao)
	if r.Descricao != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

type UpdateTipoItemRequest struct {
	Descricao string `json:"descricao"`
	TipoRefer uint   `json:"id_tp"`
}

func (r *UpdateTipoItemRequest) Validate() error {
	Logger.Debugf("Request Update: %v", r.Descricao)
	if r.Descricao != "" && r.TipoRefer > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
