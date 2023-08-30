package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/tipo_item"
)

func initializeRouterTipoItem(g *gin.RouterGroup) {
	g.POST("/tipo_item", tipo_item.CreateTipoHandler)
	g.DELETE("/tipo_item", tipo_item.DeleteTipoHandler)
	g.GET("/tipo_item", tipo_item.ShowTipoItemHandler)
	g.PUT("/tipo_item", tipo_item.UpdateipoHandler)
	g.GET("/tipo_items", tipo_item.ListTiposHandler)
}
