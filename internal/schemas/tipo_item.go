package schemas

import (
	"time"

	"gorm.io/gorm"
)

type TipoItem struct {
	gorm.Model
	Descricao string
	TipoID    uint
	Tipo      Tipo
}

type TipoItemResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeleteAt  time.Time `json:"deleteAt,omitempty"`
	TipoID    uint      `json:"id_tp"`
	Descricao string    `json:"descricao"`
}
