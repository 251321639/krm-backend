package node

import (
	"krm-backend/controllers/node"
	// "krm-backend/routers/node"

	"github.com/gin-gonic/gin"
)

//	func add(nodeGroup *gin.RouterGroup) {
//		nodeGroup.POST("/create", node.Create)
//	}
func update(nodeGroup *gin.RouterGroup) {
	nodeGroup.POST("/update", node.Update)
}

//	func delete(nodeGroup *gin.RouterGroup) {
//		nodeGroup.POST("/delete", node.Delete)
//	}
//
//	func deleteList(nodeGroup *gin.RouterGroup) {
//		nodeGroup.POST("/deleteList", node.DeleteList)
//	}
func get(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/get", node.Get)
}
func list(nodeGroup *gin.RouterGroup) {
	nodeGroup.GET("/list", node.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	nodeGroup := g.Group("/node")
	// add(nodeGroup)
	update(nodeGroup)
	// delete(nodeGroup)
	get(nodeGroup)
	list(nodeGroup)
	// deleteList(nodeGroup)
}
