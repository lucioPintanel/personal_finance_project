package tipo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Create tipo
// @Description	Save tipo data in Db.
// @Tags		tipo
// @Accept		application/json
// @Produce		application/json
// @Param		request body handler.CreateTipoRequest true "Create body"
// @Success		200 {object} handler.CreateTipoResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/tipos [post]
func CreateTipoHandler(ctx *gin.Context) {
	request := handler.CreateTipoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tipo := schemas.Tipo{
		Descricao: request.Descricao,
	}

	if err := handler.Db.Create(&tipo).Error; err != nil {
		handler.Logger.Errorf("Error creating: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "Error creating tipo on database")
		return
	}
	handler.SendSuccess(ctx, "create-tipo", tipo.ID)
}
