package pod

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

func Get(r *gin.Context) {
	logs.Debug(nil, "获取Pod详情")
	// clientset, basicInfo, err := controllers.BasicInit(r, nil)
	// returnData := config.ReturnData{}
	// if err != nil {
	// 	returnData.Message = err.Error()
	// 	returnData.Status = 400
	// 	r.JSON(200, returnData)
	// 	return
	// }
	// pod, err := clientset.CoreV1().Pods(basicInfo.Namespace).Get(context.TODO(), basicInfo.Name, metav1.GetOptions{})
	// if err != nil {
	// 	msg := "获取Pod详情失败" + err.Error()
	// 	returnData.Message = msg
	// 	returnData.Status = 400
	// } else {
	// 	returnData.Message = "获取成功"
	// 	returnData.Status = 200
	// 	returnData.Data = make(map[string]interface{})
	// 	returnData.Data["item"] = pod
	// }
	// r.JSON(200, returnData)

	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "获取成功")
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.Get(r, kubeUtilser)
}
