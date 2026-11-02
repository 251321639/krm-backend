package configmap

import (
	"krm-backend/controllers/configmap"
	// "krm-backend/routers/configmap"

	"github.com/gin-gonic/gin"
)

func add(configmapGroup *gin.RouterGroup) {
	configmapGroup.POST("/create", configmap.Create)
}
func update(configmapGroup *gin.RouterGroup) {
	configmapGroup.POST("/update", configmap.Update)
}
func delete(configmapGroup *gin.RouterGroup) {
	configmapGroup.POST("/delete", configmap.Delete)
}
func deleteList(configmapGroup *gin.RouterGroup) {
	configmapGroup.POST("/deleteList", configmap.DeleteList)
}
func get(configmapGroup *gin.RouterGroup) {
	configmapGroup.GET("/get", configmap.Get)
}
func list(configmapGroup *gin.RouterGroup) {
	configmapGroup.GET("/list", configmap.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	configmapGroup := g.Group("/configmap")
	add(configmapGroup)
	update(configmapGroup)
	delete(configmapGroup)
	get(configmapGroup)
	list(configmapGroup)
	deleteList(configmapGroup)
}
