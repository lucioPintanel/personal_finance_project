package tipo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Show tipo
// @Description	Show a tipo.
// @Tags		tipo
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Tipo identification"
// @Success		200 {object} showTipoResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Router		/tipos/{id} [get]
func ShowTipoHandler(ctx *gin.Context) {
	id := ctx.Param("id")
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
	handler.SendSuccess(ctx, "show-tipo", http.StatusOK, tipo)
}
