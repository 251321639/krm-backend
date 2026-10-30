package routers

import (
	"krm-backend/routers/auth"
	"krm-backend/routers/cluster"
	"krm-backend/routers/namespace"
	"krm-backend/routers/pod"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRoutes(apiGroup)
	cluster.RegisterSubRoutes(apiGroup)
	namespace.RegisterSubRoutes(apiGroup)
	pod.RegisterSubRouter(apiGroup)
}
