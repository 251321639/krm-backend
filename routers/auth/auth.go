package auth

import (
	"krm-backend/controllers/auth"

	"github.com/gin-gonic/gin"
)

func login(authGroup *gin.RouterGroup) {
	// 处理登录逻辑
	authGroup.POST("/login", auth.Login)
}

func logout(authGroup *gin.RouterGroup) {
	// 处理登出逻辑
	authGroup.GET("/logout", auth.Logout)
}

func RegisterSubRoutes(g *gin.RouterGroup) {
	authGroup := g.Group("/auth")
	login(authGroup)
	logout(authGroup)
}
