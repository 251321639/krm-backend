package deployment

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

func Get(r *gin.Context) {
	logs.Debug(nil, "获取deployment详情")

	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "获取成功")
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDeployment(kubeconfig, nil)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.Get(r, kubeUtilser)
}
