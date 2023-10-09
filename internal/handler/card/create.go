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
// @Param		request body createCardRequest true "Create body"
// @Success		201 {object} createCardResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		422 {object} handler.ErrorResponse
// @Router		/cards [post]
func CreateCardHandler(ctx *gin.Context) {
	request := createCardRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	card := schemas.Card{
		FlagCard: request.FlagCard,
		NumbCard: request.NumbCard,
		IDUser:   request.IDUser,
		Validity: request.Validity,
		DueDate:  request.DueDate,
	}

	if err := handler.Db.Create(&card).Error; err != nil {
		handler.Logger.Errorf("Error creating: %v", err.Error())
		handler.SendError(ctx, http.StatusUnprocessableEntity, "Error creating card on database")
		return
	}
	handler.SendSuccess(ctx, "create-card", http.StatusCreated, card.ID)
}
