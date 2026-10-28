package cluster

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(r *gin.Context) {
	logs.Info(nil, "获取集群列表")
	// listOptions := metav1.ListOptions{
	// 	,
	// }
	secretList, err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), metav1.ListOptions{LabelSelector: "kubeasy.com/cluster.metadata=true"})
	returnData := config.NewReturnData()
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "查询集群列表失败")
		// 查询失败
		msg := "查询失败：" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		r.JSON(200, returnData)
		return
	}
	// 优化数据返回的结构
	var clusterList []map[string]string
	for _, v := range secretList.Items {
		anno := v.Annotations
		clusterList = append(clusterList, anno)
	}
	returnData.Status = 200
	returnData.Message = "查询成功"
	returnData.Data = make(map[string]any)
	returnData.Data["list"] = clusterList
	r.JSON(200, returnData)
}
