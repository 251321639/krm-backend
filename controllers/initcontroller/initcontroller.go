package initcontroller

import "krm-backend/utils/logs"

func init() {
	metaDataInit()
	logs.Info(nil, "初始化基本数据")
}
