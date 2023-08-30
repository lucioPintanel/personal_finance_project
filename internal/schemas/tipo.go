package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Tipo struct {
	gorm.Model
	Descricao string
}

type TipoResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeleteAt  time.Time `json:"deleteAt,omitempty"`
	Descricao string    `json:"descricao"`
}
