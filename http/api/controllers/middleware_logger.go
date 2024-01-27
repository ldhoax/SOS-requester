package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// middlewarMiddlewareLoggereLogger logs the incoming HTTP request and response. Enable it only for debug purpose disable it on production.
func (s *Helper) MiddlewareLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		if c.FullPath() == "/v1/health" {
			// Call the next handler don't log if it is internal request from health check of Kubernetes
			c.Next()
			return
		}
		latency := time.Since(t)
		logMessage := fmt.Sprintf("path:%s, method: %s, requestBody: %v", c.FullPath(), c.Request.Method, c.Request.Body)
		s.logger.Infof("%s, duration: %v", logMessage, latency)
	}
}
