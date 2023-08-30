package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/card"
)

func initializeRouterCard(g *gin.RouterGroup) {
	g.GET("/card", card.ShowCardHandler)
	g.POST("/card", card.CreateCardHandler)
	g.PUT("/card", card.UpdateCardHandler)
	g.DELETE("/card", card.DeleteCardHandler)
	g.GET("/cards", card.ListCardsHandler)
}
