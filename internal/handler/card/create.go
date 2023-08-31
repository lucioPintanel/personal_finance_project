package card

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Create card
// @Description	Save card data in Db.
// @Tags		card
// @Accept		application/json
// @Produce		application/json
// @Param		request body handler.CreateCardRequest true "Create body"
// @Success		200 {object} handler.CreateCardRequest
// @Failure		400 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/cards [post]
func CreateCardHandler(ctx *gin.Context) {
	request := handler.CreateCardRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	card := schemas.Card{
		Flag:   request.Flag,
		Number: request.Number,
		UserId: request.UserId,
	}

	if err := handler.Db.Create(&card).Error; err != nil {
		handler.Logger.Errorf("Error creating: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "Error creating card on database")
		return
	}
	handler.SendSuccess(ctx, "create-card", card.ID)
}
