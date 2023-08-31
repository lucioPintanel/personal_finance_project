package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Show user
// @Description	Show a user.
// @Tags		user
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "User identification"
// @Success		200 {object} handler.ShowUserResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Router		/users/{id} [get]
func ShowUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		handler.SendError(ctx, http.StatusBadRequest,
			handler.ErrParamIsRequired("id",
				"queryParameter").
				Error())
		return
	}

	user := schemas.User{}

	if err := handler.Db.First(&user, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("user with id: [%s] not found", id))
		return
	}
	handler.SendSuccess(ctx, "show-user", user)
}
