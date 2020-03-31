/*
 * Package middleware
 * File: log_request.go
 * Project: SXYGo
 * File Created: 2020-03-31 15:1:31
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-31 18:43:54
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package middleware

import (
	"bytes"
	"fmt"
	"lwgo/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		startTime := time.Now()
		c.Next()

		// responseBody := bodyLogWriter.body.String()

		// var responseCode int
		// var responseMsg string
		// var responseData interface{}

		// if responseBody != "" {
		// 	response := transformers.Response{}
		// 	err := json.Unmarshal([]byte(responseBody), &response)
		// 	if err == nil {
		// 		responseCode = response.Code
		// 		responseMsg = response.Message
		// 		responseData = response.Data
		// 	}
		// }

		endTime := time.Now()

		if c.Request.Method == "POST" {
			c.Request.ParseForm()
		}

		statusCode := c.Writer.Status()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		clientIP := c.ClientIP()
		clientUA := c.Request.UserAgent()
		reqID := c.GetString("RequestID")

		reqFields := logrus.Fields{
			"req":     fmt.Sprintf("%s %s %3d", reqMethod, reqURI, statusCode),
			"ip":      reqID,
			"latency": latencyTime,
		}
		c.Set("RequestFields", reqFields)

		utils.RequestLog(c).Infof("IP: %s  UA: %s",
			clientIP,
			clientUA,
		)
	}
}
