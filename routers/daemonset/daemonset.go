package daemonset

import (
	"krm-backend/controllers/daemonset"
	// "krm-backend/routers/daemonset"

	"github.com/gin-gonic/gin"
)

func add(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.POST("/create", daemonset.Create)
}
func update(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.POST("/update", daemonset.Update)
}
func delete(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.POST("/delete", daemonset.Delete)
}
func deleteList(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.POST("/deleteList", daemonset.DeleteList)
}
func get(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.GET("/get", daemonset.Get)
}

//	func restart(daemonsetGroup *gin.RouterGroup) {
//		daemonsetGroup.GET("/restart", daemonset.Restart)
//	}
func list(daemonsetGroup *gin.RouterGroup) {
	daemonsetGroup.GET("/list", daemonset.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	daemonsetGroup := g.Group("/daemonset")
	add(daemonsetGroup)
	update(daemonsetGroup)
	delete(daemonsetGroup)
	get(daemonsetGroup)
	list(daemonsetGroup)
	deleteList(daemonsetGroup)
	// restart(daemonsetGroup)
}
