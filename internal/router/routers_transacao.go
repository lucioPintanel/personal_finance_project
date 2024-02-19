package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/transacao"
)

func initializeRouterTransacao(g *gin.RouterGroup) {
	tip := g.Group("/transacaos")
	tip.GET("/:id/*action", transacao.ShowTransacaoHandler)
	tip.POST("/", transacao.CreateTransacaoHandler)
	tip.PUT("/:id/*action", transacao.UpdateTransacaoHandler)
	tip.DELETE("/:id/*action", transacao.DeleteTransacaoHandler)
	tip.GET("/", transacao.ListTransacaosHandler)
}
