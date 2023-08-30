package tipo_item

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

func CreateTipoHandler(ctx *gin.Context) {
	request := handler.CreateTipoItemRequest{}

	ctx.BindJSON(&request)

	handler.Logger.Debugf("Request: %v", request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tipoItem := schemas.TipoItem{
		Descricao: request.Descricao,
		TipoID:    request.TipoID,
	}

	if err := handler.Db.Create(&tipoItem).Error; err != nil {
		handler.Logger.Errorf("Error creating: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "Error creating tipo on database")
		return
	}
	handler.SendSuccess(ctx, "create-tipo_item", tipoItem.ID)
}
