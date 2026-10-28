package cluster

import (
	"krm-backend/controllers/cluster"

	"github.com/gin-gonic/gin"
)

func add(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/add", cluster.Add)
}

func update(clusterGroup *gin.RouterGroup) {
	clusterGroup.POST("/update", cluster.Update)
}

func delete(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/delete", cluster.Delete)
}

func get(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/get", cluster.Get)
}

func list(clusterGroup *gin.RouterGroup) {
	clusterGroup.GET("/list", cluster.List)
}

func RegisterSubRoutes(g *gin.RouterGroup) {
	clusterGroup := g.Group("/cluster")
	add(clusterGroup)
	update(clusterGroup)
	delete(clusterGroup)
	get(clusterGroup)
	list(clusterGroup)
}
