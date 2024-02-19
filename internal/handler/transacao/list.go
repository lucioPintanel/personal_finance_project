package transacao

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		List transacao
// @Description	List all transacao.
// @Tags		transacao
// @Accept		application/json
// @Produce		application/json
// @Success		200 {object} listTransacaoResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/transacaos [get]
func ListTransacaosHandler(ctx *gin.Context) {
	transacaos := []schemas.Transacao{}

	if err := handler.Db.Find(&transacaos).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			"error listing transacao")
		return
	}
	handler.SendSuccess(ctx, "list-transacao", http.StatusOK, transacaos)
}
