package tipo_item

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Show tipo_item
// @Description	Show a tipo_item.
// @Tags		tipo_item
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Tipo Item identification"
// @Success		200 {object} handler.ShowTipoItemResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Router		/tipo_items/{id} [get]
func ShowTipoItemHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest,
			handler.ErrParamIsRequired("id",
				"queryParameter").
				Error())
		return
	}

	tipo_item := schemas.TipoItem{}

	if err := handler.Db.First(&tipo_item, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("tipo_item with id: [%s] not found", id))
		return
	}
	handler.SendSuccess(ctx, "show-tipo_item", http.StatusOK, tipo_item)
}
