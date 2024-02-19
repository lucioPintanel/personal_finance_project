package schemas

import (
	"time"

	"gorm.io/gorm"
)

type CreditCardSpending struct {
	gorm.Model
	IDCard       uint
	IDTpCategory uint
	PaymentAt    time.Time
	Descricao    string
	Valor        float64
}

type CreditCardSpendingResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeleteAt     time.Time `json:"deleteAt,omitempty"`
	IDCard       uint      `json:"id_card"`
	IDTpCategory uint      `json:"id_tp_category"`
	Payment      time.Time `json:"paymentAt"`
	Descricao    string    `json:"descricao"`
	Valor        float64   `json:"valor"`
}
