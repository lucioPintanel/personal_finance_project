package card

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		List card
// @Description	List all card.
// @Tags		card
// @Accept		application/json
// @Produce		application/json
// @Success		200 {object} handler.ListCardsResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/cards [get]
func ListCardsHandler(ctx *gin.Context) {
	cards := []schemas.Card{}

	if err := handler.Db.Find(&cards).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			"error listing card")
		return
	}
	handler.SendSuccess(ctx, "list-card", cards)
}
