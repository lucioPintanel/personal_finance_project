package card

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Delete card
// @Description	Delete card data in Db.
// @Tags		card
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Card identification"
// @Success		200 {object} handler.DeleteCardResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/card/{id} [delete]
func DeleteCardHandler(ctx *gin.Context) {
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

	if err := handler.Db.Delete(&card).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting card with id: [%s]", id))
		return
	}
	handler.SendSuccess(ctx, "delete-card", card.ID)
}
