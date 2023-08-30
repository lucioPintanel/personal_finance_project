package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nome  string
	Senha string
	Cpf   string `gorm:"size:11;unique"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeleteAt  time.Time `json:"deleteAt,omitempty"`
	Nome      string    `json:"nome"`
	Senha     string    `json:"senha"`
	Cpf       string    `json:"cpf"`
}
