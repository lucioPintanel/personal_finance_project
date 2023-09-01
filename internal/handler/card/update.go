package card

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Update card
// @Description	Update card data in Db.
// @Tags		card
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Card identification"
// @Param		request body handler.UpdateCardRequest true "Card data to Update body"
// @Success		200 {object} handler.UpdateCardResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		409 {object} handler.ErrorResponse
// @Router		/cards/{id} [put]
func UpdateCardHandler(ctx *gin.Context) {
	request := handler.UpdateCardRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Flag != "" {
		card.Flag = request.Flag
	}
	if request.Number != "" {
		card.Number = request.Number
	}
	if request.UserId > 0 {
		card.UserId = request.UserId
	}

	if err := handler.Db.Save(&card).Error; err != nil {
		handler.Logger.Errorf("error updating card: %s", err.Error())
		handler.SendError(ctx, http.StatusConflict, "error updating card")
		return
	}
	handler.SendSuccess(ctx, "update-card", http.StatusOK, card.ID)
}
