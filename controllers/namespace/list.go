package namespace

import (
	"context"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(r *gin.Context) {
	logs.Debug(nil, "查看Namespace列表")
	clientset, _, err := controllers.BasicInit(r, nil)
	returnData := config.ReturnData{}
	if err != nil {
		returnData.Message = err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		msg := "查询Namespace失败" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
	} else {
		returnData.Message = "查询成功"
		returnData.Status = 200
		returnData.Data = make(map[string]interface{})
		returnData.Data["items"] = namespaceList.Items
	}
	r.JSON(200, returnData)
}
