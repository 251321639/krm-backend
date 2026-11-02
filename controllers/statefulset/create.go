package statefulset

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
)

func Create(r *gin.Context) {
	logs.Debug(nil, "创建statefulset")
	var statefulset appsv1.StatefulSet
	//需要看一下是否要自动创建headless service
	var info controllers.Info
	info.Item = &statefulset
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	//判断是否需要自动创建Service
	if info.AutoCreateService {
		err := controllers.CreateServiceByController(
			statefulset.Spec.Template.Spec.Containers,
			statefulset.Spec.Selector.MatchLabels,
			statefulset.Namespace,
			statefulset.Name,
			kubeconfig,
			"StatefulSet")
		if err != nil {
			logs.Warning(map[string]interface{}{"错误": err.Error()}, "创建service失败")
		}
	}

	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewStatefulSet(kubeconfig, &statefulset)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.Create(r, kubeUtilser)

}
