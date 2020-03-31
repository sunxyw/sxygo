/*
 * Package middleware
 * File: request_id.go
 * Project: SXYGo
 * File Created: 2020-03-31 18:20:06
 * Author: sunxyw <xy2496419818@gmail.com>
 * -----
 * Last Modified: 2020-03-31 18:22:41
 * Modified By: sunxyw <xy2496419818@gmail.com>
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// RequestID 获取或生成请求ID
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		uRequestID := c.Request.Header.Get("X-Request-ID")

		// Create request id with UUID4
		if uRequestID == "" {
			uuid4 := uuid.NewV4()
			uRequestID = uuid4.String()
		}

		// Expose it for use in the application
		c.Set("RequestID", uRequestID)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-ID", uRequestID)
		c.Next()
	}
}
