package tipo_item

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Delete tipo_item
// @Description	Delete tipo_item data in Db.
// @Tags		tipo_item
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "User identification"
// @Success		200 {object} handler.DeleteTipoItemResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/tipo_item/{id} [delete]
func DeleteTipoHandler(ctx *gin.Context) {
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

	if err := handler.Db.Delete(&tipo_item).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting tipo_item with id: [%s]", id))
		return
	}
	handler.SendSuccess(ctx, "delete-tipo_item", tipo_item.ID)
}