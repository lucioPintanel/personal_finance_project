package transacao

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Create transacao
// @Description	Save transacao data in Db.
// @Tags		transacao
// @Accept		application/json
// @Produce		application/json
// @Param		request body createTransacaoRequest true "Create body"
// @Success		201 {object} createTransacaoResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		422 {object} handler.ErrorResponse
// @Router		/transacaos [post]
func CreateTransacaoHandler(ctx *gin.Context) {
	request := createTransacaoRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transacao := schemas.Transacao{
		IDConta:      request.IDConta,
		IDTpMoviment: request.IDTpMoviment,
		IDTpCategory: request.IDTpCategory,
		PaymentAt:    request.Payment,
		Descricao:    request.Descricao,
		Valor:        request.Valor,
	}

	if err := handler.Db.Create(&transacao).Error; err != nil {
		handler.Logger.Errorf("Error creating: %v", err.Error())
		handler.SendError(ctx, http.StatusUnprocessableEntity, "Error creating transacao on database")
		return
	}
	handler.SendSuccess(ctx, "create-transacao", http.StatusCreated, transacao.ID)
}
