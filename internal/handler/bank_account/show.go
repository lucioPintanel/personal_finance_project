package bank_account

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Show bank_account
// @Description	Show a bank_account.
// @Tags		bank_account
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "BankAccount identification"
// @Success		200 {object} showBankAccountResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Router		/bank_accounts/{id} [get]
func ShowBankAccountHandler(ctx *gin.Context) {
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
	handler.SendSuccess(ctx, "show-bankAccount", http.StatusOK, bankAccount)
}
