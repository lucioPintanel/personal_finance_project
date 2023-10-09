package schemas

import (
	"time"

	"gorm.io/gorm"
)

type BankAccount struct {
	gorm.Model
	IDUser    uint
	Descricao string
	NroConta  string
	Balance   float64
}

type BankAccountResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeleteAt  time.Time `json:"deleteAt,omitempty"`
	IDUser    uint      `json:"id_user"`
	Descricao string    `json:"descricao"`
	NroConta  string    `json:"nro_conta"`
	Balance   float64   `json:"balance"`
}
