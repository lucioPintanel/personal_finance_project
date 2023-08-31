package tipo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		List tipo
// @Description	List all tipo.
// @Tags		tipo
// @Accept		application/json
// @Produce		application/json
// @Success		200 {object} handler.ListTipoResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/tipo [get]
func ListTiposHandler(ctx *gin.Context) {
	tipos := []schemas.Tipo{}

	if err := handler.Db.Find(&tipos).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			"error listing tipo")
		return
	}
	handler.SendSuccess(ctx, "list-tipo", tipos)
}
