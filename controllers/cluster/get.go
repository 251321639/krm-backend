package cluster

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(r *gin.Context) {
	logs.Info(nil, "获取集群详情")
	clusterId := r.Query("clusterId")
	returnData := config.ReturnData{}
	clusterSecret, err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Get(context.TODO(), clusterId, metav1.GetOptions{})
	if err != nil {
		logs.Error(map[string]any{"msg": err.Error()}, "获取集群信息详情")
		returnData.Status = 400
		returnData.Message = "获取集群信息详情"
		r.JSON(200, returnData)
		return
	}
	returnData.Status = 200
	returnData.Message = "查询成功"
	returnData.Data = make(map[string]any)
	clusterConfigMap := clusterSecret.Annotations
	clusterConfigMap["kubeconfig"] = string(clusterSecret.Data["kubeconfig"])
	returnData.Data["item"] = clusterConfigMap
	r.JSON(200, returnData)
}
