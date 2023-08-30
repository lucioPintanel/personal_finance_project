package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	Flag   string
	Number string `gorm:"size:16;unique"`
	UserId uint
}

type CardResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeleteAt  time.Time `json:"deleteAt,omitempty"`
	Flag      string    `json:"flag"`
	Number    string    `json:"number"`
	UserId    uint      `json:"user_id"`
}
