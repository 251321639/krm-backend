package cluster

import (
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

func Update(r *gin.Context) {
	logs.Debug(nil, "更新集群")
	addOrUpdate(r, "Update")
}
