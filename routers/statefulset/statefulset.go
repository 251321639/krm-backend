package statefulset

import (
	"krm-backend/controllers/statefulset"
	// "krm-backend/routers/statefulset"

	"github.com/gin-gonic/gin"
)

func add(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.POST("/create", statefulset.Create)
}
func update(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.POST("/update", statefulset.Update)
}
func delete(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.POST("/delete", statefulset.Delete)
}
func deleteList(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.POST("/deleteList", statefulset.DeleteList)
}
func get(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.GET("/get", statefulset.Get)
}

//	func restart(statefulsetGroup *gin.RouterGroup) {
//		statefulsetGroup.GET("/restart", statefulset.Restart)
//	}
func list(statefulsetGroup *gin.RouterGroup) {
	statefulsetGroup.GET("/list", statefulset.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	statefulsetGroup := g.Group("/statefulset")
	add(statefulsetGroup)
	update(statefulsetGroup)
	delete(statefulsetGroup)
	get(statefulsetGroup)
	list(statefulsetGroup)
	deleteList(statefulsetGroup)
	// restart(statefulsetGroup)
}
