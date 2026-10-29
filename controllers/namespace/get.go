package namespace

import (
	"context"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(r *gin.Context) {
	logs.Debug(nil, "获取namespace详情")
	clientset, basicInfo, err := controllers.BasicInit(r, nil)
	returnData := config.ReturnData{}
	if err != nil {
		returnData.Message = err.Error()
		returnData.Status = 400
		r.JSON(200, returnData)
		return
	}
	namespace, err := clientset.CoreV1().Namespaces().Get(context.TODO(), basicInfo.Name, metav1.GetOptions{})
	if err != nil {
		msg := "获取Namespace详情失败" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
	} else {
		returnData.Message = "获取成功"
		returnData.Status = 200
		returnData.Data = make(map[string]interface{})
		returnData.Data["item"] = namespace
	}
	r.JSON(200, returnData)
}
