package transacao

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Delete transacap
// @Description	Delete transacap data in Db.
// @Tags		transacap
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Transacao identification"
// @Success		204 {object} deleteTransacaoResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/transacaps/{id} [delete]
func DeleteTransacaoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest,
			handler.ErrParamIsRequired("id",
				"queryParameter").
				Error())
		return
	}

	transacap := schemas.Transacao{}

	if err := handler.Db.First(&transacap, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("transacap with id: [%s] not found", id))
		return
	}

	if err := handler.Db.Delete(&transacap).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting transacap with id: [%s]", id))
		return
	}
	handler.SendSuccess(ctx, "delete-transacap", http.StatusNoContent, transacap.ID)
}
