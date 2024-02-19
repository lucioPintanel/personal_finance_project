package transacao

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Show transacao
// @Description	Show a transacao.
// @Tags		transacao
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Transacao identification"
// @Success		200 {object} showTransacaoResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Router		/transacaos/{id} [get]
func ShowTransacaoHandler(ctx *gin.Context) {
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
	handler.SendSuccess(ctx, "show-transacao", http.StatusOK, transacao)
}
