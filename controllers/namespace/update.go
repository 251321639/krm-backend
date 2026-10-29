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

func Update(r *gin.Context) {
	logs.Debug(nil, "更新namespace")
	var ns corev1.Namespace
	clientset, _, err := controllers.BasicInit(r, &ns)
	returnData := config.ReturnData{}
	if err != nil {
		returnData.Message = err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}

	_, err = clientset.CoreV1().Namespaces().Update(context.TODO(), &ns, metav1.UpdateOptions{})
	if err != nil {
		msg := "更新namespace失败" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
	} else {
		returnData.Message = "更新成功"
		returnData.Status = 200
	}
	r.JSON(200, returnData)
}
