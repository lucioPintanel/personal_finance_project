package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/tipo_item"
)

func initializeRouterTipoItem(g *gin.RouterGroup) {
	i := g.Group("/tipo_items")
	i.GET("/:id/*action", tipo_item.ShowTipoItemHandler)
	i.POST("/", tipo_item.CreateTipoItemHandler)
	i.DELETE("/:id/*action", tipo_item.DeleteTipoItemHandler)
	i.PUT("/:id/*action", tipo_item.UpdateTipoItemHandler)
	i.GET("/", tipo_item.ListTiposItemHandler)
}
