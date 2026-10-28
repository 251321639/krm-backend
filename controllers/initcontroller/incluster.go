package initcontroller

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils/logs"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func metaDataInit() {
	logs.Info(nil, "初始化元数据命名空间")
	restConfig, err := clientcmd.BuildConfigFromFlags("", "config/meta.kubeconfig")
	if err != nil {
		logs.Error(map[string]any{"msg": err.Error()}, "inCluster kubeconfig加载失败")
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		logs.Error(map[string]any{"msg": err.Error()}, "inCluster客户端创建失败")
		panic(err.Error())
	}
	config.InClusterClientSet = clientset
	inClusterInfo, _ := clientset.Discovery().ServerVersion()
	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), config.MetadataNamespace, metav1.GetOptions{})
	if err != nil {
		var metadataNamespace corev1.Namespace
		metadataNamespace.Name = config.MetadataNamespace
		_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &metadataNamespace, metav1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]any{"msg": err.Error()}, "元数据命名空间创建失败")
			panic(err.Error())
		}

		logs.Info(map[string]any{"Namespace": config.MetadataNamespace, "inCluster版本": inClusterInfo.GitVersion}, "元数据命名空间创建成功")
	} else {
		logs.Info(map[string]any{"Namespace": config.MetadataNamespace, "inCluster版本": inClusterInfo.GitVersion}, "元数据命名空间已存在")
	}

}
