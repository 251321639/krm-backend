package cluster

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils/logs"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/gin-gonic/gin"
)

func Delete(r *gin.Context) {
	logs.Info(nil, "删除集群")
	// 1.接收删除集群的ID
	clusterId := r.Query("clusterId")
	// 2.删除
	err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Delete(context.TODO(), clusterId, metav1.DeleteOptions{})
	returnData := config.NewReturnData()
	if err != nil {
		logs.Error(map[string]any{"clusterId": clusterId}, "集群删除失败")
		msg := "集群删除失败：" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
	} else {
		logs.Warning(map[string]any{"clusterId": clusterId}, "集群删除成功")
		returnData.Status = 200
		returnData.Message = "删除成功"
		delete(config.ClusterKubeconfig, clusterId)
	}
	r.JSON(200, returnData)
}
