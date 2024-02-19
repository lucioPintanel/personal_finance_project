package transacao

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Update transacao
// @Description	Update transacao data in Db.
// @Tags		transacao
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Transacao identification"
// @Param		request body updateTransacaoRequest true "Transacao data to Update body"
// @Success		200 {object} updateTransacaoResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/transacaos/{id} [put]
func UpdateTransacaoHandler(ctx *gin.Context) {
	request := updateTransacaoRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
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

	transacao := schemas.Transacao{}

	if err := handler.Db.First(&transacao, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("transacao with id: [%s] not found", id))
		return
	}

	if request.Descricao != "" {
		transacao.Descricao = request.Descricao
	}

	if err := handler.Db.Save(&transacao).Error; err != nil {
		handler.Logger.Errorf("error updating transacao: %s", err.Error())
		handler.SendError(ctx, http.StatusConflict, "error updating transacao")
		return
	}
	handler.SendSuccess(ctx, "update-transacao", http.StatusOK, transacao.ID)
}
