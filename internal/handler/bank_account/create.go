package bank_account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Create bank_account
// @Description	Save bank_account data in Db.
// @Tags		bank_account
// @Accept		application/json
// @Produce		application/json
// @Param		request body createBankAccountRequest true "Create body"
// @Success		201 {object} createBankAccountResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		422 {object} handler.ErrorResponse
// @Router		/bank_accounts [post]
func CreateBankAccountHandler(ctx *gin.Context) {
	request := createBankAccountRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		handler.Logger.Errorf("Validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	bankAccount := schemas.BankAccount{
		IDUser:    request.IDUser,
		Descricao: request.Descricao,
		NroConta:  request.NroConta,
		Balance:   request.Balance,
	}

	if err := handler.Db.Create(&bankAccount).Error; err != nil {
		handler.Logger.Errorf("Error creating: %v", err.Error())
		handler.SendError(ctx, http.StatusUnprocessableEntity, "Error creating bankAccount on database")
		return
	}
	handler.SendSuccess(ctx, "create-bankAccount", http.StatusCreated, bankAccount.ID)
}
