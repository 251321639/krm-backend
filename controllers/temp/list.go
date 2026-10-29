package temp

import (
	// "context"
	// "krm-backend/config"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(r *gin.Context) {
	logs.Info(nil, "获取Namespace列表")

}
