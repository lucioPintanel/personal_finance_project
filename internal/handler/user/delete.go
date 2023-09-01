package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Delete user
// @Description	Delete user data in Db.
// @Tags		user
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "User identification"
// @Success		204 {object} deleteUserResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/users/{id} [delete]
func DeleteUserHandler(ctx *gin.Context) {
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

	if err := handler.Db.Delete(&user).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting user with id: [%s]", id))
		return
	}
	handler.SendSuccess(ctx, "delete-user", http.StatusNoContent, user.ID)
}
