package web

import (
	"github.com/gin-gonic/gin"
	"go-memo/pkg/e"
	"go-memo/pkg/util"
	"time"
)

// @Author KHighness
// @Update 2022-02-13

// json web token
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = 200
		var data interface{}
		token := c.GetHeader("Authorization")

		// Token验证失败直接返回错误信息
		if token == "" {
			code = e.ErrorAuthCheckTokenFail
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenExpire
			}
		}
		if code != e.SUCCESS {
			c.JSON(400, gin.H{
				"status" : code,
				"msg"    : e.GetMsg(code),
				"data"   : data,
			})
			c.Abort()
			return
		}

		// 验证成功进入下一步处理
		c.Next()
	}
}