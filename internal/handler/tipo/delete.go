package tipo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Delete tipo
// @Description	Delete tipo data in Db.
// @Tags		tipo
// @Accept		application/json
// @Produce		application/json
// @Param		id query string true "Tipo identification"
// @Success		200 {object} handler.DeleteTipoResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/tipo [delete]
func DeleteTipoHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest,
			handler.ErrParamIsRequired("id",
				"queryParameter").
				Error())
		return
	}

	tipo := schemas.Tipo{}

	if err := handler.Db.First(&tipo, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("tipo with id: [%s] not found", id))
		return
	}

	if err := handler.Db.Delete(&tipo).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting tipo with id: [%s]", id))
		return
	}
	handler.SendSuccess(ctx, "delete-tipo", tipo.ID)
}