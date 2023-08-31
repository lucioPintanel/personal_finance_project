package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/card"
)

func initializeRouterCard(g *gin.RouterGroup) {
	c := g.Group("/cards")
	c.GET("/:id/*action", card.ShowCardHandler)
	c.POST("/", card.CreateCardHandler)
	c.PUT("/:id/*action", card.UpdateCardHandler)
	c.DELETE("/:id/*action", card.DeleteCardHandler)
	c.GET("/", card.ListCardsHandler)
}
