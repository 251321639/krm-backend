package cluster

import (
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Context) {
	logs.Debug(nil, "添加集群")
	addOrUpdate(r, "Create")
}
