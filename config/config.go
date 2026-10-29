// 配置层 存放程序的配置信息
package config

import (
	"krm-backend/utils/logs"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
)

const (
	TimeFormat                    string = "2006-01-02 15:04:05"
	ClusterConfigSecretLabelKey   string = "kubeasy.com/cluster.metadata"
	ClusterConfigSecretLabelValue string = "true"
)

var (
	Port string
	//jwt
	JwtSignKey string
	JwtExpTime int64 //jwt token过期时间，单位，分钟
	Username   string
	Password   string
	//inCluster相关配置
	MetadataNamespace string //元数据存储的namespace
	//inCluster 客户端
	InClusterClientSet *kubernetes.Clientset
	// ClusterKubeconfig   map[string]string
	RunAsMode           string
	ForbidNamespaceList []string
	ClusterKubeconfig   map[string]string

	InCluster string
)

// var ClusterKubeconfig = make(map[string]string)

type ReturnData struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

// 构造函数，用来自定义初始化结构体里面的值
func NewReturnData() ReturnData {
	returnData := ReturnData{}
	returnData.Status = 200
	data := make(map[string]any)

	returnData.Data = data
	return returnData
}

func initLogConfig(logLevel string) {
	//配置程序的日志输出级别
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	//文件名和行号加进去
	logrus.SetReportCaller(true)
	//日志格式改成json
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TimeFormat,
		//runtime.Frame: 帧，可用于获取调用者返回的PC值的函数，文件或者行信息
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := path.Base(f.File)
			return f.Function, fileName
		},
	})
}

func init() {
	logs.Debug(nil, "开始加载程序配置")
	//环境变量加载我们的配置
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("GIN_MODE", "debug")
	//获取程序启动端口号配置,电脑未设置环境变量默认未8080
	viper.SetDefault("PORT", ":8080")
	//获取jwt加密的secret
	viper.SetDefault("JWT_SIGN_KEY", "try")
	//获取jwt过期时间配置
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	//配置用户名密码的默认值，我系统中有一个username::***是一个坑，一定要检查
	//加密用户名和密码 md5
	viper.SetDefault("user_name", "admin")
	viper.SetDefault("password", "1234")
	//inCluster相关配置
	viper.SetDefault("METADATA_NAMESPACE", "krm")
	//禁止删除某些namespace，可以通过环境变量进行配置
	viper.SetDefault("FORBID_NAMESPACES", "krm,kube-system,default,kube-public,kube-public")

	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL") //获取程序的设置
	RunAsMode = viper.GetString("GIN_MODE")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpTime = viper.GetInt64("JWT_EXPIRE_TIME")
	//获取用户名的密码
	Username = viper.GetString("user_name")
	Password = viper.GetString("password")
	//inCluster相关配置
	MetadataNamespace = viper.GetString("METADATA_NAMESPACE")
	//test1,test2,test3
	forbidnamespaces := viper.GetString("FORBID_NAMESPACES")
	ForbidNamespaceList = strings.Split(forbidnamespaces, ",")

	//加载日志输出格式
	initLogConfig(logLevel)

	//inCluster相关配置
	viper.SetDefault("IN_CLUSTER", "false")
	InCluster = viper.GetString("IN_CLUSTER")
}
