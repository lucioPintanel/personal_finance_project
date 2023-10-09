package bank_account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		List bank_account
// @Description	List all bank_account.
// @Tags		bank_account
// @Accept		application/json
// @Produce		application/json
// @Success		200 {object} listBankAccountsResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/bank_accounts [get]
func ListBankAccountsHandler(ctx *gin.Context) {
	bankAccounts := []schemas.BankAccount{}

	if err := handler.Db.Find(&bankAccounts).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			"error listing bank_account")
		return
	}
	handler.SendSuccess(ctx, "list-bank_account", http.StatusOK, bankAccounts)
}
