package auth

import (
	"krm-backend/config"
	"krm-backend/utils/jwtutil"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(r *gin.Context) {
	userInfo := UserInfo{}
	returnData := config.NewReturnData()
	if err := r.ShouldBindJSON(&userInfo); err != nil {
		returnData.Status = 401
		returnData.Message = err.Error()
		r.JSON(401, returnData)
		return
	}
	logs.Debug(map[string]any{"用户名": userInfo.Username, "密码": userInfo.Password}, "用户登录请求")

	// 这里可以添加实际的登录逻辑，例如验证用户名和密码
	if userInfo.Username == config.Username && userInfo.Password == config.Password {
		token, err := jwtutil.GenToken(userInfo.Username)
		if err != nil {
			logs.Error(map[string]any{"用户名": userInfo.Username}, "生成token失败")
			returnData.Status = 500
			returnData.Message = "生成token失败"
			returnData.Data = nil
			r.JSON(500, returnData)
			return
		}
		logs.Debug(map[string]any{"用户名": userInfo.Username, "token": token}, "用户登录成功，生成token")
		returnData.Status = 200
		returnData.Message = "登录成功"
		returnData.Data = map[string]any{"token": token}
		r.JSON(200, returnData)

		// r.Set("Authorization", token)
	} else {
		logs.Warning(map[string]any{"用户名": userInfo.Username}, "用户名或密码错误")
		returnData.Status = 401
		returnData.Message = "用户名或密码错误"
		returnData.Data = nil
		r.JSON(401, returnData)
	}
}

func Logout(r *gin.Context) {
	logs.Debug(nil, "用户登出请求")
	returnData := config.NewReturnData()
	returnData.Status = 200
	returnData.Message = "登出成功"
	returnData.Data = nil
	// 这里可以添加实际的登出逻辑，例如清除会话或令牌
	r.JSON(200, returnData)
}
