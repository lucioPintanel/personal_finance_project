package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/tipo_item"
)

func initializeRouterTipoItem(g *gin.RouterGroup) {
	g.POST("/tipo_item", tipo_item.CreateTipoHandler)
}
