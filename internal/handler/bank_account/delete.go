package bank_account

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

// @BasePath  /api/v1

// @Summary		Delete bank_account
// @Description	Delete bank_account data in Db.
// @Tags		bank_account
// @Accept		application/json
// @Produce		application/json
// @Param		id path int true "BankAccount identification"
// @Success		204 {object} deleteBankAccountResponse
// @Failure		400 {object} handler.ErrorResponse
// @Failure		404 {object} handler.ErrorResponse
// @Failure		500 {object} handler.ErrorResponse
// @Router		/bank_accounts/{id} [delete]
func DeleteBankAccountHandler(ctx *gin.Context) {
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

	if err := handler.Db.Delete(&bankAccount).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError,
			fmt.Sprintf("error deleting bankAccount with id: [%s]", id))
		return
	}
	handler.SendSuccess(ctx, "delete-bankAccount", http.StatusNoContent, bankAccount.ID)
}
