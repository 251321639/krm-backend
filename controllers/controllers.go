// 控制器层 实现路由的处理逻辑
package controllers

import (
	"errors"
	"krm-backend/config"
	"krm-backend/utils/logs"
	"strconv"

	// "github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义一个map用于存储每个集群的资源数量
var ClusterStatistics map[string]map[string]int

// 定义一个控制factory关闭的map
var FactoryStopper map[string]chan struct{}

// 定义全局的数据结构
type BasicInfo struct {
	ClusterId  string      `json:"clusterId" form:"clusterId"`
	Namespace  string      `json:"namespace" form:"namespace"`
	Name       string      `json:"name" form:"name"`
	Item       interface{} `json:"item"`
	DeleteList []string    `json:"deleteList"`
}

type Info struct {
	BasicInfo
	ReturnData    config.ReturnData
	LabelSelector string `json:"labelSelector" form:"labelSelector"`
	FieldSelector string `json:"fieldSelector" form:"fieldSelector"`
	//判断是否是强制删除
	Force bool `json:"force" form:"force"`
	// namespace 复制字段
	DestClusterId     string              `json:"destClusterId" form:"destClusterId"`
	DestNamespace     string              `json:"destNamespace" form:"destNamespace"`
	CreateNamespace   bool                `json:"createNamespace" form:"createNamespace"`
	Resources         map[string][]string `json:"resources" form:"resources"`
	AutoCreateService bool                `json:"autoCreateService" form:"autoCreateService"`
	Yaml              string              `json:"yaml" form:"yaml"`
	Method            string              `json:"method" form:"method"`
	// Apply             bool                `json:"apply" form:"apply"`
	Container      string `json:"container" form:"container"`
	TailLines      int64  `json:"tailLines" form:"tailLines"`
	DefaultCommand string `json:"defaultCommand" form:"defaultCommand"`
}

func NewInfo(r *gin.Context, info *Info, returnDataMsg string) (kubeconfig string) {
	//首先获取请求的类型
	requestMethod := r.Request.Method
	var err error
	// var returnData config.ReturnData
	info.ReturnData.Message = returnDataMsg
	info.ReturnData.Status = 200
	if requestMethod == "GET" {
		err = r.ShouldBindQuery(&info)
	} else if requestMethod == "POST" {
		err = r.ShouldBindBodyWithJSON(&info)
	} else {
		err = errors.New("不支持的请求类型")
	}
	if err != nil {
		msg := "请求出错:" + err.Error()
		info.ReturnData.Message = msg
		info.ReturnData.Status = 400
		logs.Error(nil, msg)
		r.JSON(200, info.ReturnData)
		return
	}
	//获取kubeconfig
	kubeconfig = config.ClusterKubeconfig[info.ClusterId]
	return kubeconfig
}

// 自动添加service后端逻辑
func CreateServiceByController(containers []corev1.Container, labels map[string]string, namespace, name, kubeconfig, resourceType string) (err error) {
	logs.Debug(map[string]interface{}{"容器": containers, "标签": labels, "命名空间": namespace, "名称": name, "资源类型": resourceType}, "自动创建service")
	//创建一个Service对象
	var service corev1.Service
	service.Name = name
	service.Spec.Selector = make(map[string]string)
	service.Labels = make(map[string]string)
	service.Annotations = make(map[string]string)
	service.Labels = labels
	service.Annotations["kubeasy.com/autoCreate"] = "true"
	service.Spec.Selector = labels
	//处理service端口
	//拿到容器的配置
	for _, c := range containers {
		//拿到容器的端口配置
		for portIndex, containerPort := range c.Ports {
			var servicePort corev1.ServicePort
			if containerPort.Name == "" {
				//处理多个Port 名称不能为空
				servicePort.Name = c.Name + "-" + strconv.Itoa(portIndex)
			} else {
				servicePort.Name = containerPort.Name
			}
			servicePort.Protocol = containerPort.Protocol
			servicePort.Port = containerPort.ContainerPort
			service.Spec.Ports = append(service.Spec.Ports, servicePort)
		}
	}
	//判断是否是statefuleset
	if resourceType == "StatefulSet" {
		service.Spec.ClusterIP = "None"
	}
	return
	//创建Service
	// instance := kubeutils.NewService(kubeconfig, &service)
	// err = instance.Create(namespace)
	// return
}

// func (c *Info) Create(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
// 	err := kubeUtilsInterface.Create(c.Namespace)
// 	if err != nil {
// 		msg := "创建失败" + err.Error()
// 		c.ReturnData.Message = msg
// 		c.ReturnData.Status = 400
// 		logs.Error(nil, msg)
// 	}
// 	r.JSON(200, c.ReturnData)
// }

// func (c *Info) Update(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
// 	err := kubeUtilsInterface.Update(c.Namespace)
// 	if err != nil {
// 		msg := "更新失败" + err.Error()
// 		c.ReturnData.Message = msg
// 		c.ReturnData.Status = 400
// 		logs.Error(nil, msg)
// 	}
// 	r.JSON(200, c.ReturnData)
// }

// func (c *Info) List(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
// 	items, err := kubeUtilsInterface.List(c.Namespace, c.LabelSelector, c.FieldSelector)
// 	if err != nil {
// 		msg := "查询失败" + err.Error()
// 		c.ReturnData.Message = msg
// 		c.ReturnData.Status = 400
// 		logs.Error(nil, msg)
// 	} else {
// 		c.ReturnData.Data = make(map[string]interface{})
// 		c.ReturnData.Data["items"] = items
// 	}
// 	r.JSON(200, c.ReturnData)
// }

// func (c *Info) Get(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
// 	item, err := kubeUtilsInterface.Get(c.Namespace, c.Name)
// 	if err != nil {
// 		msg := "查询失败" + err.Error()
// 		c.ReturnData.Message = msg
// 		c.ReturnData.Status = 400
// 		logs.Error(nil, msg)
// 	} else {
// 		c.ReturnData.Data = make(map[string]interface{})
// 		c.ReturnData.Data["item"] = item
// 	}
// 	r.JSON(200, c.ReturnData)
// }

// func (c *Info) Delete(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
// 	var gracePeriodSeconds int64
// 	if c.Force {
// 		//强制删除
// 		var s int64 = 0
// 		gracePeriodSeconds = s
// 	}
// 	err := kubeUtilsInterface.Delete(c.Namespace, c.Name, &gracePeriodSeconds)
// 	if err != nil {
// 		msg := "删除失败" + err.Error()
// 		c.ReturnData.Message = msg
// 		c.ReturnData.Status = 400
// 		logs.Error(nil, msg)
// 	}
// 	r.JSON(200, c.ReturnData)
// }

// func (c *Info) DeleteList(r *gin.Context, kubeUtilsInterface kubeutils.KubeUtilser) {
// 	var gracePeriodSeconds int64
// 	if c.Force {
// 		//强制删除
// 		var s int64 = 0
// 		gracePeriodSeconds = s
// 	}
// 	err := kubeUtilsInterface.DeleteList(c.Namespace, c.BasicInfo.DeleteList, &gracePeriodSeconds)
// 	if err != nil {
// 		msg := "删除失败" + err.Error()
// 		c.ReturnData.Message = msg
// 		c.ReturnData.Status = 400
// 		logs.Error(nil, msg)
// 	}
// 	r.JSON(200, c.ReturnData)
// }

func BasicInit(r *gin.Context, item interface{}) (clientset *kubernetes.Clientset, basicInfo BasicInfo, err error) {
	basicInfo = BasicInfo{}
	basicInfo.Item = item
	//首先获取请求的类型
	switch r.Request.Method {
	case "GET":
		err = r.ShouldBindQuery(&basicInfo)
	case "POST":
		err = r.ShouldBindBodyWithJSON(&basicInfo)
	default:
		err = errors.New("不支持的请求类型")
	}
	if err != nil {
		msg := "请求出错:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	if basicInfo.Namespace == "" {
		basicInfo.Namespace = "default"
	}
	//获取kubeconfig
	kubeconfig := config.ClusterKubeconfig[basicInfo.ClusterId]
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		msg := "解析kubeconfig失败:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	clientset, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "创建clientset失败:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	return clientset, basicInfo, nil
}
