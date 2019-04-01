package middleware

import (
	"fmt"
	"github.com/plimble/ace"
	"time"
)

// APIStatsD 统计api请求
func APIStatsD() ace.HandlerFunc {
	return func(c *ace.C) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"),c.Request.URL,c.Request.Method,c.Request.RemoteAddr)
		c.Next()
	}
}
