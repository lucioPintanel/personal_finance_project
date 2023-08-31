package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/tipo"
)

func initializeRouterTipo(g *gin.RouterGroup) {
	tip := g.Group("/tipos")
	tip.GET("/:id/*action", tipo.ShowTipoHandler)
	tip.POST("/", tipo.CreateTipoHandler)
	tip.PUT("/:id/*action", tipo.UpdateipoHandler)
	tip.DELETE("/:id/*action", tipo.DeleteTipoHandler)
	tip.GET("/", tipo.ListTiposHandler)
}
