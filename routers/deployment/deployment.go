package deployment

import (
	"krm-backend/controllers/deployment"
	// "krm-backend/routers/deployment"

	"github.com/gin-gonic/gin"
)

func add(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.POST("/create", deployment.Create)
}
func update(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.POST("/update", deployment.Update)
}
func delete(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.POST("/delete", deployment.Delete)
}
func deleteList(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.POST("/deleteList", deployment.DeleteList)
}
func get(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.GET("/get", deployment.Get)
}
func restart(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.GET("/restart", deployment.Restart)
}
func list(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.GET("/list", deployment.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	deploymentGroup := g.Group("/deployment")
	add(deploymentGroup)
	update(deploymentGroup)
	delete(deploymentGroup)
	get(deploymentGroup)
	list(deploymentGroup)
	deleteList(deploymentGroup)
	restart(deploymentGroup)
}
