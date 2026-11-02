package pv

import (
	"krm-backend/controllers/pv"
	// "krm-backend/routers/pv"

	"github.com/gin-gonic/gin"
)

func add(pvGroup *gin.RouterGroup) {
	pvGroup.POST("/create", pv.Create)
}
func update(pvGroup *gin.RouterGroup) {
	pvGroup.POST("/update", pv.Update)
}
func delete(pvGroup *gin.RouterGroup) {
	pvGroup.POST("/delete", pv.Delete)
}
func deleteList(pvGroup *gin.RouterGroup) {
	pvGroup.POST("/deleteList", pv.DeleteList)
}
func get(pvGroup *gin.RouterGroup) {
	pvGroup.GET("/get", pv.Get)
}
func list(pvGroup *gin.RouterGroup) {
	pvGroup.GET("/list", pv.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	pvGroup := g.Group("/pv")
	add(pvGroup)
	update(pvGroup)
	delete(pvGroup)
	get(pvGroup)
	list(pvGroup)
	deleteList(pvGroup)
}
