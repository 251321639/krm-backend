package namespace

import (
	"context"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create(r *gin.Context) {
	logs.Debug(nil, "添加namespace")
	clientset, basicInfo, err := controllers.BasicInit(r, nil)
	returnData := config.ReturnData{}
	if err != nil {
		returnData.Message = err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	var namespace corev1.Namespace
	namespace.Name = basicInfo.Name
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
	if err != nil {
		msg := "创建Namespace失败" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
	} else {
		returnData.Message = "创建成功"
		returnData.Status = 200
	}
	r.JSON(200, returnData)
}
