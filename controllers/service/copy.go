package service

import (
	"krm-backend/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	corev1 "k8s.io/api/core/v1"
)

func Copy(srckubeconfig, destkubeconfig, srcNamespace, destNamespace string, list []string) {
	var service corev1.Service
	//创建源集群deployment实例，用来查询需要复制的数据配置
	srcInstance := kubeutils.NewService(srckubeconfig, &service)
	//创建目标集群deployment实例，用来创建需要复制的数据配置
	destInstance := kubeutils.NewService(destkubeconfig, &service)
	//第一步，拿到所有需要复制的数据
	for _, name := range list {
		logs.Info(map[string]interface{}{"资源类型": "service",
			"资源名称": name, "目标命名空间": destNamespace}, "开始拷贝数据")
		//在源集群当中查询该数据
		_, err := srcInstance.Get(srcNamespace, name)
		if err != nil {
			logs.Error(map[string]interface{}{"资源类型": "service",
				"资源名称": name, "目标命名空间": destNamespace, "msg": err.Error()}, "拷贝失败,源数据查询失败")
			continue
		}
		//说明源集群有这个资源
		destInstance.Item = srcInstance.Item
		destInstance.Item.Namespace = destNamespace
		destInstance.Item.ResourceVersion = "" // 清除 resourceVersion
		err = destInstance.Create(destNamespace)
		if err != nil {
			logs.Error(map[string]interface{}{"资源类型": "service",
				"资源名称": name, "目标命名空间": destNamespace, "msg": err.Error()}, "拷贝失败,目标集群创建失败")

		}
	}
}
