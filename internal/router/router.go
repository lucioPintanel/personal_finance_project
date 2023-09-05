package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/configs"
	"github.com/lucioPintanel/personal_finance_project/internal/handler"

	docs "github.com/lucioPintanel/personal_finance_project/docs"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Initialize() {
	//Inicializando logger e banco de dados
	handler.InitializeHandler()

	r := gin.Default()
	//basePath := "/api/v1"
	basePath := configs.GetServerApiPath() + configs.GetServerVersion()
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)

	//Inicializando as rotas da API
	initializeRouterTipo(v1)
	initializeRouterTipoItem(v1)
	initializeRouterUser(v1)
	initializeRouterCard(v1)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	port := configs.GetServerPort()
	r.Run(":" + port)
}
