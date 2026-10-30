package pod

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

func Create(r *gin.Context) {
	logs.Debug(nil, "创建pod")
	var pod corev1.Pod
	var info controllers.Info
	info.Item = &pod
	kubeconfig := controllers.NewInfo(r, &info, "创建成功")
	// clientset, basicInfo, err := controllers.BasicInit(r, &pod)
	// returnData := config.ReturnData{}
	// if err != nil {
	// 	returnData.Message = err.Error()
	// 	returnData.Status = 400
	// 	r.JSON(200, returnData)
	// 	return
	// }
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, &pod)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.Create(r, kubeUtilser)
	// _, err = clientset.CoreV1().Pods(basicInfo.Namespace).Create(context.TODO(), &pod, metav1.CreateOptions{})
	// if err != nil {
	// 	msg := "创建Pod失败" + err.Error()
	// 	returnData.Message = msg
	// 	returnData.Status = 400
	// } else {
	// 	returnData.Message = "创建成功"
	// 	returnData.Status = 200
	// }
	// r.JSON(200, returnData)
}
