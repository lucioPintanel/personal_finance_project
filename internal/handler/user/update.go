package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Update user
// @Description	Update user data in Db.
// @Tags		user
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "User identification"
// @Param		request body handler.UpdateUserRequest true "User data to Update body"
// @Success		200 {object} handler.UpdateCardResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/users/{id} [put]
func UpdateUserHandler(ctx *gin.Context) {
	request := handler.UpdateUserRequest{}

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

	user := schemas.User{}

	if err := handler.Db.First(&user, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("user with id: [%s] not found", id))
		return
	}

	if request.Nome != "" {
		user.Nome = request.Nome
	}
	if request.Senha != "" {
		user.Senha = request.Senha
	}

	if err := handler.Db.Save(&user).Error; err != nil {
		handler.Logger.Errorf("error updating user: %s", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}
	handler.SendSuccess(ctx, "update-user", user.ID)
}
