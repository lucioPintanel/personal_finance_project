package card

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type CreateCardResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type DeleteCardResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type ShowCardResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type UpdateCardResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type ListCardsResponse struct {
	handler.BaseResponse
	Data []schemas.CardResponse `json:"data"`
}
