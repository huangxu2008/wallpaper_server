package middlewares

import (
	"net/http"
	"strings"
	"wallpaper_server/utils"

	"github.com/gin-gonic/gin"
)

// Token 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 解析 Token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 将用户 ID 存入上下文
		c.Set("userID", userID)
		c.Next()
	}
}
