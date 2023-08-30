package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucioPintanel/personal_finance_project/internal/handler/user"
)

func initializeRouterUser(g *gin.RouterGroup) {
	g.GET("/user", user.ShowUserHandler)
	g.POST("/user", user.CreateUser1Handler)
	g.PUT("/user", user.UpdateUserHandler)
	g.DELETE("/user", user.DeleteUserHandler)
	g.GET("/users", user.ListUsersHandler)
}
