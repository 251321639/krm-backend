package daemonset

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
)

func Update(r *gin.Context) {
	logs.Debug(nil, "更新daemonset")
	var daemonset appsv1.DaemonSet
	var info controllers.Info
	info.Item = &daemonset
	kubeconfig := controllers.NewInfo(r, &info, "更新成功")
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, &daemonset)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.Update(r, kubeUtilser)

}
