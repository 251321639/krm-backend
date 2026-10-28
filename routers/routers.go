package routers

import (
	"krm-backend/routers/auth"
	"krm-backend/routers/cluster"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRoutes(apiGroup)
	cluster.RegisterSubRoutes(apiGroup)
}
