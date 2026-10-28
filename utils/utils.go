// 工具类
package utils

import "encoding/json"

func Struct2Map(s any) map[string]string {
	//结构体转json再转map
	j, _ := json.Marshal(s)
	m := make(map[string]string)
	json.Unmarshal(j, &m)
	return m
}
