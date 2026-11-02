package daemonset

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
)

func Create(r *gin.Context) {
	logs.Debug(nil, "创建daemonset")
	var daemonset appsv1.DaemonSet
	//需要看一下是否要自动创建headless service
	var info controllers.Info
	info.Item = &daemonset
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	//判断是否需要自动创建Service
	if info.AutoCreateService {
		//创建service
		// controllers.CreateServiceByController(deployment.Spec.Template.Spec.Containers, deployment.Spec.Selector.MatchLabels, deployment.Namespace, deployment.Name, kubeconfig, "deployment")
		//创建service
		err := controllers.CreateServiceByController(
			daemonset.Spec.Template.Spec.Containers,
			daemonset.Spec.Selector.MatchLabels,
			daemonset.Namespace,
			daemonset.Name,
			kubeconfig,
			"DaemonSet")
		if err != nil {
			logs.Warning(map[string]interface{}{"错误": err.Error()}, "创建service失败")
		}
	}
	// clientset, basicInfo, err := controllers.BasicInit(r, &deployment)
	// returnData := config.ReturnData{}
	// if err != nil {
	// 	returnData.Message = err.Error()
	// 	returnData.Status = 400
	// 	r.JSON(200, returnData)
	// 	return
	// }
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewDaemonSet(kubeconfig, &daemonset)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.Create(r, kubeUtilser)

}
