package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/tipo"
)

func initializeRouterTipo(g *gin.RouterGroup) {
	g.GET("/tipo", tipo.ShowTipoHandler)
	g.POST("/tipo", tipo.CreateTipoHandler)
	g.PUT("/tipo", tipo.UpdateipoHandler)
	g.DELETE("/tipo", tipo.DeleteTipoHandler)
	g.GET("/tipos", tipo.ListTiposHandler)
}
