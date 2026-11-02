package pv

//如何对集群中的namespace进行管理
//1.首先要知道对哪个集群中的namespace进行管理
//2.kubeconfig，secret，clusterId --> secret --> data --> kubeconfig --> clientset
//	kubeconfig存储在一个变量当中,map[string]string
//3.创建一个clientset
//4.操作namespace
//
