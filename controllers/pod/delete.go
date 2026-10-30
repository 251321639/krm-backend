package pod

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

func Delete(r *gin.Context) {
	logs.Debug(nil, "删除pod")
	// clientset, basicInfo, err := controllers.BasicInit(r, nil)
	// returnData := config.ReturnData{}
	// if err != nil {
	// 	returnData.Message = err.Error()
	// 	returnData.Status = 400
	// 	r.JSON(200, returnData)
	// 	return
	// }
	// for _, v := range basicInfo.DeleteList {
	// 	_ = clientset.CoreV1().Pods(basicInfo.Namespace).Delete(context.TODO(), v, metav1.DeleteOptions{})
	// }

	// returnData.Message = "删除成功"
	// returnData.Status = 200
	// r.JSON(200, returnData)
	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.Delete(r, kubeUtilser)
}

func DeleteList(r *gin.Context) {
	logs.Debug(nil, "删除pod")
	// clientset, basicInfo, err := controllers.BasicInit(r, nil)
	// returnData := config.ReturnData{}
	// if err != nil {
	// 	returnData.Message = err.Error()
	// 	returnData.Status = 400
	// 	r.JSON(200, returnData)
	// 	return
	// }
	// for _, v := range basicInfo.DeleteList {
	// 	_ = clientset.CoreV1().Pods(basicInfo.Namespace).Delete(context.TODO(), v, metav1.DeleteOptions{})
	// }

	// returnData.Message = "删除成功"
	// returnData.Status = 200
	// r.JSON(200, returnData)
	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "删除成功")
	//改为接口形式的创建
	var kubeUtilser kubeutils.KubeUtilser
	instance := kubeutils.NewPod(kubeconfig, nil)
	//把实例复制给kubeUtilser
	kubeUtilser = instance
	//使用kubeUtilser创建资源
	// kubeUtilser.Create("default")
	info.DeleteList(r, kubeUtilser)
}
