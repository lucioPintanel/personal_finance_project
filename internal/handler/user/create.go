package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Create user
// @Description	Save user data in Db.
// @Tags		user
// @Accept		application/json
// @Produce		application/json
// @Param		request body handler.CreateUserRequest true "Create body"
// @Success		200 {object} handler.CreateUserRequest
// @Failure		400 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/user [post]
func CreateUser1Handler(ctx *gin.Context) {
	request := handler.CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Nome:  request.Nome,
		Senha: request.Senha,
		Cpf:   request.Cpf,
	}

	if err := handler.Db.Create(&user).Error; err != nil {
		handler.Logger.Errorf("Error creating: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "Error creating user on database")
		return
	}
	handler.SendSuccess(ctx, "create-user", user.ID)
}