package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/bank_account"
)

func initializeRouterBankAccount(g *gin.RouterGroup) {
	c := g.Group("/bank_accounts")
	c.GET("/:id/*action", bank_account.ShowBankAccountHandler)
	c.POST("/", bank_account.CreateBankAccountHandler)
	c.PUT("/:id/*action", bank_account.UpdateBankAccountHandler)
	c.DELETE("/:id/*action", bank_account.DeleteBankAccountHandler)
	c.GET("/", bank_account.ListBankAccountsHandler)
}
