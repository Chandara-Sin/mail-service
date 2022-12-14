package logger

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const key = "logger"

func Middleware(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("X-API-KEY")
		l := log.With(zap.String("x-api-id", apiKey))
		c.Set(key, l)

		bodyBytes := []byte{}
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		fmt.Printf("request body: %s\n", bodyBytes)

		c.Next()
	}
}

func Unwrap(c *gin.Context) *zap.Logger {
	val, _ := c.Get(key)
	if log, ok := val.(*zap.Logger); ok {
		return log
	}
	return zap.NewExample()
}
