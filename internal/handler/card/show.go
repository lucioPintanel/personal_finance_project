package card

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Show card
// @Description	Show a card.
// @Tags		card
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Card identification"
// @Success		200 {object} showCardResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Router		/cards/{id} [get]
func ShowCardHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest,
			handler.ErrParamIsRequired("id",
				"queryParameter").
				Error())
		return
	}

	card := schemas.Card{}

	if err := handler.Db.First(&card, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("card with id: [%s] not found", id))
		return
	}
	handler.SendSuccess(ctx, "show-card", http.StatusOK, card)
}
