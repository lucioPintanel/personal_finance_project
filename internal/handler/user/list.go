package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		List user
// @Description	List all user.
// @Tags		user
// @Accept		application/json
// @Produce		application/json
// @Success		200 {object} handler.ListUsersResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/user [get]
func ListUsersHandler(ctx *gin.Context) {
	users := []schemas.User{}

	if err := handler.Db.Find(&users).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			"error listing user")
		return
	}
	handler.SendSuccess(ctx, "list-user", users)
}
