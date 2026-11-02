package ingress

import (
	"krm-backend/controllers"
	"krm-backend/utils/logs"
	"net/http"
	"strconv"
	"strings"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

//	rules:
//
// - host: kubeasy.com
// http:
//
//	  paths:
//		- path: /
//		  pathType: Prefix
//		  backend:
//			service:
//			  name: nginx
//			  port:
//				number: 80
//
// 1.先获取当前ingress配置
// 2.找到ingress所以的rule
// 3.找到所有的域名配置
// 4.匹配每一个域名配置，找到所有的path
// 5.找到每一个path对应的service
// 6.找到每一个service对应的pod
// 定义Node结构体
type Node struct {
	Id           string `json:"id"`
	Label        string `json:"label"`
	ResourceType string `json:"resourceType"`
}

// 定义Edge结构体
type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func Topology(r *gin.Context) {
	logs.Debug(nil, "获取ingress拓扑图")
	var info controllers.Info
	kubeconfig := controllers.NewInfo(r, &info, "获取成功")
	// 打印 namespace 和 name，便于调试
	logs.Debug(map[string]interface{}{
		"namespace": info.Namespace,
		"name":      info.Name,
	}, "请求的 Ingress 信息")
	instance := kubeutils.NewIngress(kubeconfig, nil)
	serviceInstance := kubeutils.NewService(kubeconfig, nil)
	podInstance := kubeutils.NewPod(kubeconfig, nil)
	test, err := instance.Get(info.Namespace, info.Name)

	if err != nil {
		logs.Error(map[string]interface{}{"msg": err.Error(),
			"命名空间": info.Namespace,
			"名称":   info.Name,
		}, "获取ingress配置失败")
		info.ReturnData.Message = "获取拓扑图失败:" + err.Error()
		info.ReturnData.Status = 400
		r.JSON(http.StatusOK, info.ReturnData)
		return
	}

	// 打印获取到的 Ingress 对象，便于调试
	if test == nil {
		logs.Debug(map[string]interface{}{
			"namespace": info.Namespace,
			"name":      info.Name,
		}, "获取到的 Ingress 对象详情")
	} else {
		logs.Debug(map[string]interface{}{
			"ingress": test,
		}, "获取到的 Ingress 对象详情")
	}

	// 检查 instance.Item 是否为 nil
	if instance.Item == nil {
		logs.Error(map[string]interface{}{
			"命名空间": info.Namespace,
			"名称":   info.Name,
		}, "获取到的 Ingress 对象为空")
		info.ReturnData.Message = "获取到的 Ingress 对象为空"
		info.ReturnData.Status = 400
		r.JSON(http.StatusOK, info.ReturnData)
		return
	}
	var nodes []Node
	var edges []Edge
	var podFullList []string
	//存储service关联的pod列表
	servicePod := make(map[string][]string)
	//已经处理过的域名
	hostUsed := make(map[string]string)
	//已经处理过的service
	serviceUsed := make(map[string]string)
	//已经处理过service和pod有哪些
	serviceIdPodUsed := make(map[string]string)
	ingress := instance.Item

	//拿到所有的rules规则
	rules := ingress.Spec.Rules
	for _, rule := range rules {
		host := rule.Host
		if host == "" {
			host = "*"
		}
		if _, ok := hostUsed[host]; !ok {
			// 如果没有使用过这个host，就添加到map中
			hostNode := Node{
				Id:           host,
				Label:        host,
				ResourceType: "domain",
			}
			nodes = append(nodes, hostNode)
			hostUsed[host] = "true"
		}
		for _, path := range rule.HTTP.Paths {
			pathId := host + path.Path + string(*path.PathType)
			pathNode := Node{
				Id:           pathId,
				Label:        path.Path,
				ResourceType: "path",
			}
			nodes = append(nodes, pathNode)
			hostPathEdge := Edge{
				From: host,
				To:   pathId,
			}
			edges = append(edges, hostPathEdge)
			backendService := path.Backend.Service
			serviceName := backendService.Name
			portString := ""
			if backendService.Port.Number == 0 {
				portString = backendService.Port.Name
			} else {
				portString = strconv.Itoa(int(backendService.Port.Number))
			}
			serviceId := serviceName + ":" + portString
			if _, ok := serviceUsed[serviceId]; !ok {
				// 如果没有使用过这个service，就添加到map中
				serviceNode := Node{
					Id:           serviceId,
					Label:        serviceId,
					ResourceType: "service",
				}
				nodes = append(nodes, serviceNode)
				serviceUsed[serviceId] = "true"
			}
			//链接path和service
			pathServiceEdge := Edge{
				From: pathId,
				To:   serviceId,
			}
			edges = append(edges, pathServiceEdge)
			//首先获取到当前service配置
			if _, ok := servicePod[serviceName]; !ok {
				_, err := serviceInstance.Get(info.Namespace, serviceName)
				var servicePodList []string
				if err == nil {
					//map
					var selectorList []string
					selector := serviceInstance.Item.Spec.Selector
					for k, v := range selector {
						s := k + "=" + v
						selectorList = append(selectorList, s)
					}
					//查询pod
					_, err := podInstance.List(info.Namespace, strings.Join(selectorList, ","), "")
					if err == nil {
						for _, pod := range podInstance.Items {
							//记录当前service关联的pod列表
							servicePodList = append(servicePodList, pod.Name)
							podExist := false
							for _, name := range podFullList {
								if name == pod.Name {
									//说明pod列表中当前已经有pod，说明nodes当中有当前的pod的node了
									podExist = true
									break
								}
							}
							if podExist {
								continue
							}
							//把当前pod存储到nodes
							podNode := Node{
								Id:           pod.Name,
								Label:        pod.Name,
								ResourceType: "pod",
							}
							nodes = append(nodes, podNode)
							podFullList = append(podFullList, pod.Name)
						}
					}
				}
				//保存当前service和pod列表的映射关系
				servicePod[serviceName] = servicePodList
			}
			//把service连接到pod
			if _, ok := serviceIdPodUsed[serviceId]; !ok {
				//先拿到当前service关联的pod，然后把每个Pod连接到该Service上
				for _, podName := range servicePod[serviceName] {
					podEdge := Edge{
						From: serviceId,
						To:   podName,
					}
					edges = append(edges, podEdge)
				}
				serviceIdPodUsed[serviceId] = "true"
			}
		}
	}
	info.ReturnData.Data = make(map[string]interface{})
	info.ReturnData.Data["nodes"] = nodes
	info.ReturnData.Data["edges"] = edges
	r.JSON(http.StatusOK, info.ReturnData)
}
