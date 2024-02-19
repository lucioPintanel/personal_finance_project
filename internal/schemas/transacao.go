package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Transacao struct {
	gorm.Model
	IDConta      uint
	IDTpMoviment uint
	IDTpCategory uint
	PaymentAt    time.Time
	Descricao    string
	Valor        float64
}

type TransacaoResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeleteAt     time.Time `json:"deleteAt,omitempty"`
	IDConta      uint      `json:"id_conta"`
	IDTpMoviment uint      `json:"id_tp_mov"`
	IDTpCategory uint      `json:"id_tp_category"`
	Payment      time.Time `json:"paymentAt"`
	Descricao    string    `json:"descricao"`
	Valor        float64   `json:"valor"`
}
