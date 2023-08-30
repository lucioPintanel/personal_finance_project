package tipo_item

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Update tipo_item
// @Description	Update tipo_item data in Db.
// @Tags		tipo_item
// @Accept		application/json
// @Produce		application/json
// @Param		id query string true "Tipo Item identification"
// @Param		request body handler.UpdateTipoItemRequest true "Tipo data to Update body"
// @Success		200 {object} handler.UpdateTipoItemResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/tipo_item [put]
func UpdateipoHandler(ctx *gin.Context) {
	request := handler.UpdateTipoItemRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
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

	if request.Descricao != "" {
		tipo_item.Descricao = request.Descricao
	}

	if err := handler.Db.Save(&tipo_item).Error; err != nil {
		handler.Logger.Errorf("error updating tipo_item: %s", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error updating tipo_item")
		return
	}
	handler.SendSuccess(ctx, "update-tipo_item", tipo_item.ID)
}
