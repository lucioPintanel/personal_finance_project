package tipo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Update tipo
// @Description	Update tipo data in Db.
// @Tags		tipo
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "Tipo identification"
// @Param		request body handler.UpdateTipoRequest true "Tipo data to Update body"
// @Success		200 {object} handler.UpdateTipoResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/tipos/{id} [put]
func UpdateipoHandler(ctx *gin.Context) {
	request := handler.UpdateTipoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
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

	tipo := schemas.Tipo{}

	if err := handler.Db.First(&tipo, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("tipo with id: [%s] not found", id))
		return
	}

	if request.Descricao != "" {
		tipo.Descricao = request.Descricao
	}

	if err := handler.Db.Save(&tipo).Error; err != nil {
		handler.Logger.Errorf("error updating tipo: %s", err.Error())
		handler.SendError(ctx, http.StatusConflict, "error updating tipo")
		return
	}
	handler.SendSuccess(ctx, "update-tipo", http.StatusOK, tipo.ID)
}
