package temp

import (
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

func Create(r *gin.Context) {
	logs.Debug(nil, "添加Namespace")
	addOrUpdate(r, "Create")
}
