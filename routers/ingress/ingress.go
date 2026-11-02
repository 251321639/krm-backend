package ingress

import (
	"krm-backend/controllers/ingress"
	// "krm-backend/routers/ingress"

	"github.com/gin-gonic/gin"
)

func add(ingressGroup *gin.RouterGroup) {
	ingressGroup.POST("/create", ingress.Create)
}
func update(ingressGroup *gin.RouterGroup) {
	ingressGroup.POST("/update", ingress.Update)
}
func delete(ingressGroup *gin.RouterGroup) {
	ingressGroup.POST("/delete", ingress.Delete)
}
func deleteList(ingressGroup *gin.RouterGroup) {
	ingressGroup.POST("/deleteList", ingress.DeleteList)
}
func get(ingressGroup *gin.RouterGroup) {
	ingressGroup.GET("/get", ingress.Get)
}

func topology(ingressGroup *gin.RouterGroup) {
	ingressGroup.GET("/topology", ingress.Topology)
}
func list(ingressGroup *gin.RouterGroup) {
	ingressGroup.GET("/list", ingress.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	ingressGroup := g.Group("/ingress")
	add(ingressGroup)
	update(ingressGroup)
	delete(ingressGroup)
	get(ingressGroup)
	list(ingressGroup)
	deleteList(ingressGroup)
	topology(ingressGroup)
}
