package pvc

import (
	"krm-backend/controllers/pvc"
	// "krm-backend/routers/pvc"

	"github.com/gin-gonic/gin"
)

func add(pvcGroup *gin.RouterGroup) {
	pvcGroup.POST("/create", pvc.Create)
}
func update(pvcGroup *gin.RouterGroup) {
	pvcGroup.POST("/update", pvc.Update)
}
func delete(pvcGroup *gin.RouterGroup) {
	pvcGroup.POST("/delete", pvc.Delete)
}
func deleteList(pvcGroup *gin.RouterGroup) {
	pvcGroup.POST("/deleteList", pvc.DeleteList)
}
func get(pvcGroup *gin.RouterGroup) {
	pvcGroup.GET("/get", pvc.Get)
}
func list(pvcGroup *gin.RouterGroup) {
	pvcGroup.GET("/list", pvc.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	pvcGroup := g.Group("/pvc")
	add(pvcGroup)
	update(pvcGroup)
	delete(pvcGroup)
	get(pvcGroup)
	list(pvcGroup)
	deleteList(pvcGroup)
}
