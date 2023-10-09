package card

import (
	"fmt"
	"time"

	"github.com/lucioPintanel/personal_finance_project/internal/handler"
)

type createCardRequest struct {
	FlagCard string    `json:"flag"`
	NumbCard string    `json:"number"`
	IDUser   uint      `json:"user_id"`
	Validity time.Time `json:"validity"`
	DueDate  uint8     `json:"due_date"`
}

func (r *createCardRequest) validate() error {
	handler.Logger.Debugf("Request Update: %v", r.FlagCard)
	if r.FlagCard != "" && r.NumbCard != "" && r.IDUser > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}

type updateCardRequest struct {
	FlagCard string    `json:"flag"`
	NumbCard string    `json:"number"`
	IDUser   uint      `json:"user_id"`
	Validity time.Time `json:"validity"`
	DueDate  uint8     `json:"due_date"`
}

func (r *updateCardRequest) validate() error {
	handler.Logger.Debugf("Request Update: %v", r.NumbCard)
	if r.FlagCard != "" && r.NumbCard != "" && r.IDUser > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
