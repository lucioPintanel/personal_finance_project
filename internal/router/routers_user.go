package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/user"
)

func initializeRouterUser(g *gin.RouterGroup) {
	u := g.Group("/users")
	u.GET("/:id/*action", user.ShowUserHandler)
	u.POST("/", user.CreateUser1Handler)
	u.PUT("/:id/*action", user.UpdateUserHandler)
	u.DELETE("/:id/*action", user.DeleteUserHandler)
	u.GET("/", user.ListUsersHandler)
}
