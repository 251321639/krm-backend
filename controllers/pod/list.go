package pod

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

func List(r *gin.Context) {
	logs.Debug(nil, "查询Pod列表")
	// clientset, basicInfo, err := controllers.BasicInit(r, nil)
	// returnData := config.ReturnData{}
	// if err != nil {
	// 	returnData.Message = err.Error()
	// 	returnData.Status = 400
	// 	r.JSON(200, returnData)
	// 	return
	// }
	// list, err := clientset.CoreV1().Pods(basicInfo.Namespace).List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	msg := "查询pod失败" + err.Error()
	// 	returnData.Message = msg
	// 	returnData.Status = 400
	// } else {
	// 	returnData.Message = "查询成功"
	// 	returnData.Status = 200
	// 	returnData.Data = make(map[string]interface{})
	// 	returnData.Data["items"] = list.Items
	// }
	// r.JSON(200, returnData)

	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "查询成功")
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.List(r, kubeUtilser)
}
