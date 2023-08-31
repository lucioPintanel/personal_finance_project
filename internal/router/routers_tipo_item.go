package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/tipo_item"
)

func initializeRouterTipoItem(g *gin.RouterGroup) {
	i := g.Group("/tipo_item")
	i.GET("/:id/*action", tipo_item.ShowTipoItemHandler)
	i.POST("/", tipo_item.CreateTipoHandler)
	i.DELETE("/:id/*action", tipo_item.DeleteTipoHandler)
	i.PUT("/:id/*action", tipo_item.UpdateipoHandler)
	i.GET("/", tipo_item.ListTiposHandler)
}
