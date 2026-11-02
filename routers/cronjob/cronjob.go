package cronjob

import (
	"krm-backend/controllers/cronjob"
	// "krm-backend/routers/cronjob"

	"github.com/gin-gonic/gin"
)

func add(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.POST("/create", cronjob.Create)
}
func update(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.POST("/update", cronjob.Update)
}
func delete(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.POST("/delete", cronjob.Delete)
}
func deleteList(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.POST("/deleteList", cronjob.DeleteList)
}
func get(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.GET("/get", cronjob.Get)
}
func list(cronjobGroup *gin.RouterGroup) {
	cronjobGroup.GET("/list", cronjob.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	cronjobGroup := g.Group("/cronjob")
	add(cronjobGroup)
	update(cronjobGroup)
	delete(cronjobGroup)
	get(cronjobGroup)
	list(cronjobGroup)
	deleteList(cronjobGroup)
}
