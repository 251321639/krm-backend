package secret

import (
	"krm-backend/controllers/secret"
	// "krm-backend/routers/secret"

	"github.com/gin-gonic/gin"
)

func add(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/create", secret.Create)
}
func update(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/update", secret.Update)
}
func delete(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/delete", secret.Delete)
}
func deleteList(secretGroup *gin.RouterGroup) {
	secretGroup.POST("/deleteList", secret.DeleteList)
}
func get(secretGroup *gin.RouterGroup) {
	secretGroup.GET("/get", secret.Get)
}
func list(secretGroup *gin.RouterGroup) {
	secretGroup.GET("/list", secret.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	secretGroup := g.Group("/secret")
	add(secretGroup)
	update(secretGroup)
	delete(secretGroup)
	get(secretGroup)
	list(secretGroup)
	deleteList(secretGroup)
}
