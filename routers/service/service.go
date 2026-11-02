package service

import (
	"krm-backend/controllers/service"
	// "krm-backend/routers/service"

	"github.com/gin-gonic/gin"
)

func add(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/create", service.Create)
}
func update(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/update", service.Update)
}
func delete(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/delete", service.Delete)
}
func deleteList(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/deleteList", service.DeleteList)
}
func get(serviceGroup *gin.RouterGroup) {
	serviceGroup.GET("/get", service.Get)
}
func list(serviceGroup *gin.RouterGroup) {
	serviceGroup.GET("/list", service.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	serviceGroup := g.Group("/service")
	add(serviceGroup)
	update(serviceGroup)
	delete(serviceGroup)
	get(serviceGroup)
	list(serviceGroup)
	deleteList(serviceGroup)
}
