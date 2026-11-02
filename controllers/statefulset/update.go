package statefulset

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
)

func Update(r *gin.Context) {
	logs.Debug(nil, "更新StatefulSet")
	var statefulset appsv1.StatefulSet
	var info controllers.Info
	info.Item = &statefulset
	kubeconfig := controllers.NewInfo(r, &info, "更新成功")
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, &statefulset)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.Update(r, kubeUtilser)

}
