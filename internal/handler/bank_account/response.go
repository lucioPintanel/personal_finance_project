package bank_account

import (
	"github.com/lucioPintanel/personal_finance_project/internal/handler"
	"github.com/lucioPintanel/personal_finance_project/internal/schemas"
)

type createBankAccountResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type deleteBankAccountResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type showBankAccountResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type updateBankAccountResponse struct {
	handler.BaseResponse
	Data int `json:"data"`
}

type listBankAccountsResponse struct {
	handler.BaseResponse
	Data []schemas.BankAccountResponse `json:"data"`
}
