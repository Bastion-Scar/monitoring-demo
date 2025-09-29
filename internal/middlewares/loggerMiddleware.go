package middlewares

import (
	"monitoring-demo/internal/storage"
	"monitoring-demo/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		after := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		requestURI := c.Request.RequestURI
		path := c.Request.URL.Path

		logger.Info("Request Info",
			zap.Int("status", status),
			zap.String("clientIP", clientIP),
			zap.String("path", path),
			zap.String("requestURI", requestURI),
			zap.Duration("duration", after))
		if logger == nil {
			panic("logger not set")
		}

		codeRaw, codeExists := c.Get("code")
		urlRaw, urlExists := c.Get("url")

		var code, url string
		if codeStr, ok := codeRaw.(string); codeExists && ok {
			code = codeStr
		}
		if urlStr, ok := urlRaw.(string); urlExists && ok {
			url = urlStr
		}

		entry := storage.Logs{
			Duration: after,
			Status:   status,
			IP:       clientIP,
			Code:     code,
			URL:      url,
			Path:     path,
		}
		models.LogChan <- entry
	}

}
