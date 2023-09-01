package card

import (
	"fmt"

	"github.com/lucioPintanel/personal_finance_project/internal/handler"
)

type createCardRequest struct {
	Flag   string `json:"flag"`
	Number string `json:"number"`
	UserId uint   `json:"user_id"`
}

func (r *createCardRequest) validate() error {
	handler.Logger.Debugf("Request Update: %v", r.Flag)
	if r.Flag != "" && r.Number != "" && r.UserId > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

type updateCardRequest struct {
	Flag   string `json:"flag"`
	Number string `json:"number"`
	UserId uint   `json:"user_id"`
}

func (r *updateCardRequest) validate() error {
	handler.Logger.Debugf("Request Update: %v", r.Number)
	if r.Flag != "" && r.Number != "" && r.UserId > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
