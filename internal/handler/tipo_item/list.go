package tipo_item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		List tipo_item
// @Description	List all tipo_item.
// @Tags		tipo_item
// @Accept		application/json
// @Produce		application/json
// @Success		200 {object} listTipoItemsResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/tipo_items [get]
func ListTiposItemHandler(ctx *gin.Context) {
	tipo_items := []schemas.TipoItem{}

	if err := handler.Db.Find(&tipo_items).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			"error listing tipo_item")
		return
	}
	handler.SendSuccess(ctx, "list-tipo_item", http.StatusOK, tipo_items)
}
