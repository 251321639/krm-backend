package pod

import (
	"krm-backend/config"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

func Update(r *gin.Context) {
	logs.Debug(nil, "更新Pod")
	var returnData config.ReturnData
	returnData.Message = "Pod暂不支持更新操作"
	returnData.Status = 200
	r.JSON(200, returnData)

}
