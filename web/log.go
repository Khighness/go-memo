package web

import (
	"github.com/gin-gonic/gin"
	"go-memo/pkg/util"
	"time"
)

// @Author KHighness
// @Update 2022-02-13

func NewLogger() gin.HandlerFunc {
	logger := util.Logger()
	return func(c *gin.Context) {
		startTime := time.Now()               // 开始时间
		c.Next()                              // 处理请求
		endTime := time.Now()                 // 结束时间
		latencyTime := endTime.Sub(startTime) // 执行时间
		reqMethod := c.Request.Method         // 请求方式
		reqUri := c.Request.RequestURI        // 请求路由
		statusCode := c.Writer.Status()       // 状态码
		clientIp := c.ClientIP()              // 请求IP
		logger.Infof("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIp, reqMethod, reqUri)
	}
}