// package main

// import (
// 	"context"
// 	"fmt"
// 	"krm-backend/utils/logs"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// )

// func main() {
// 	config, err := clientcmd.BuildConfigFromFlags("", "./config/meta.kubeconfig")
// 	if err != nil {
// 		panic(err)
// 	}
// 	client, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err)
// 	}
// 	pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		logs.Error(nil, "查询Pod列表失败")
// 	}
// 	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

// 	deploys, err := client.AppsV1().Deployments("kube-system").List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		logs.Error(nil, "查询Deployment列表失败")
// 	}
// 	deploymentItems := deploys.Items
// 	for _, deploy := range deploymentItems {
// 		fmt.Printf("当前deploy的名字是: %s\n", deploy.Name)
// 		// logs.Info(nil, deploy.Name)
// 	}
// }
