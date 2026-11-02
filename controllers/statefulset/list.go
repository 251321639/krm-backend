package statefulset

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

func List(r *gin.Context) {
	logs.Debug(nil, "查询StatefulSet列表")

	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "查询成功")
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, nil)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.List(r, kubeUtilser)
}
