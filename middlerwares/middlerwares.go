// 中间件层
package middlewares

import (
	"krm-backend/config"
	"krm-backend/utils/jwtutil"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth(r *gin.Context) {
	//1.除了login和logout之外所有接口都要验证请求是否携带token，并且token是否合法
	requestUrl := r.FullPath()
	logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "")
	if requestUrl == "/api/auth/login" || requestUrl == "/api/auth/logout" || requestUrl == "/api/pod/logs" || requestUrl == "/api/pod/exec" {
		logs.Debug(map[string]interface{}{"请求路径": requestUrl}, "登录和退出不需要验证token")
		r.Next()
		return
	}
	//拿到returnData
	returnData := config.NewReturnData()
	//token
	//其他接口要验证token
	//获取是否携带token
	tokenString := r.Request.Header.Get("Authorization")
	if tokenString == "" {
		//说明没有携带token
		returnData.Status = 401
		returnData.Message = "请求未携带Token,请登录后尝试"
		r.JSON(200, returnData)
		r.Abort()
		return
	}
	//token不为空，验证token是否合法
	claims, err := jwtutil.ParseToken(tokenString)
	if err != nil {
		returnData.Status = 401
		returnData.Message = "Token验证未通过"
		r.JSON(200, returnData)
		r.Abort()
		return
	}
	//验证成功
	r.Set("claims", claims)
	r.Next()
}

func Cors(ctx *gin.Context) {
	logs.Debug(nil, "允许跨域")
	method := ctx.Request.Method
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	ctx.Next()
}
