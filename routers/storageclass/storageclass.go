package storageclass

import (
	"krm-backend/controllers/storageclass"
	// "krm-backend/routers/storageclass"

	"github.com/gin-gonic/gin"
)

//	func add(storageclassGroup *gin.RouterGroup) {
//		storageclassGroup.POST("/create", storageclass.Create)
//	}
//
//	func update(storageclassGroup *gin.RouterGroup) {
//		storageclassGroup.POST("/update", storageclass.Update)
//	}
//
//	func delete(storageclassGroup *gin.RouterGroup) {
//		storageclassGroup.POST("/delete", storageclass.Delete)
//	}
//
//	func deleteList(storageclassGroup *gin.RouterGroup) {
//		storageclassGroup.POST("/deleteList", storageclass.DeleteList)
//	}
//
//	func get(storageclassGroup *gin.RouterGroup) {
//		storageclassGroup.GET("/get", storageclass.Get)
//	}
func list(storageclassGroup *gin.RouterGroup) {
	storageclassGroup.GET("/list", storageclass.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	storageclassGroup := g.Group("/storageclass")
	// add(storageclassGroup)
	// update(storageclassGroup)
	// delete(storageclassGroup)
	// get(storageclassGroup)
	list(storageclassGroup)
	// deleteList(storageclassGroup)
}
