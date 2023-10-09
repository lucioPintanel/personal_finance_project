// swag init -g cmd/api-server/main.go

package main

import (
	config "github.com/lucioPintanel/personal_finance_project/configs"
	"github.com/lucioPintanel/personal_finance_project/internal/router"

	docs "github.com/lucioPintanel/personal_finance_project/docs"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email lm.pintanel@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	err := config.Init()
	if err != nil {
		//Adicionar log
		return
	}

	// programatically set swagger info
	docs.SwaggerInfo.Title = "API REST Personal Finance Project"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/v1"

	router.Initialize()
}
