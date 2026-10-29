package temp

import (
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

func Update(r *gin.Context) {
	logs.Debug(nil, "更新Namespace")
	addOrUpdate(r, "Update")
}
