package middleware

import (
	"demo04/pkg/errno"
	"demo04/pkg/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var code int32
		var data interface{}
		var errorNo errno.ErrNo
		//code = 200
		token := c.GetHeader(viper.GetString("server.jwtHeader"))
		//token := c.Query("token")
		if token == "" {
			//code = 404
			errorNo = errno.ErrorAuthCheckTokenError
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				//code = errno.ErrorAuthCheckTokenFail
				errorNo = errno.ErrorAuthCheckTokenTimeoutError
			} else if time.Now().Unix() > claims.ExpiresAt {
				//code = errno.ErrorAuthCheckTokenTimeout
				errorNo = errno.ErrorAuthCheckTokenTimeoutError
			} else {
				errorNo = errno.Success
			}
		}

		if errorNo.ErrorCode != errno.SuccessCode {
			c.JSON(200, gin.H{
				"code": errorNo.ErrorCode,
				"msg":  errorNo.ErrorMsg,
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
