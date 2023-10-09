package bank_account

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Update bank_account
// @Description	Update bank_account data in Db.
// @Tags		bank_account
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "BankAccount identification"
// @Param		request body updateBankAccountRequest true "BankAccount data to Update body"
// @Success		200 {object} updateBankAccountResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		409 {object} handler.ErrorResponse
// @Router		/bank_accounts/{id} [put]
func UpdateBankAccountHandler(ctx *gin.Context) {
	request := updateBankAccountRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
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

	bankAccount := schemas.BankAccount{}

	if err := handler.Db.First(&bankAccount, id).Error; err != nil {
		handler.SendError(ctx, http.StatusNotFound,
			fmt.Sprintf("bankAccount with id: [%s] not found", id))
		return
	}

	if request.NroConta != "" {
		bankAccount.NroConta = request.NroConta
	}
	if request.Descricao != "" {
		bankAccount.Descricao = request.Descricao
	}
	if request.IDUser > 0 {
		bankAccount.IDUser = request.IDUser
	}
	if request.Balance > 0 {
		bankAccount.Balance = request.Balance
	}

	if err := handler.Db.Save(&bankAccount).Error; err != nil {
		handler.Logger.Errorf("error updating bankAccount: %s", err.Error())
		handler.SendError(ctx, http.StatusConflict, "error updating bankAccount")
		return
	}
	handler.SendSuccess(ctx, "update-bankAccount", http.StatusOK, bankAccount.ID)
}
