package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	IDUser   uint
	FlagCard string
	NumbCard string
	Validity time.Time
	DueDate  uint8
}

type CardResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeleteAt  time.Time `json:"deleteAt,omitempty"`
	IDUser    uint      `json:"user_id"`
	FlagCard  string    `json:"flag_card"`
	NumbCard  string    `json:"numb_card"`
	Validity  time.Time `json:"validity"`
	DueDate   uint8     `json:"due_date"`
}
