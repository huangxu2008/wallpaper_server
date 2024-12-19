package middlewares

import (
	"strings"
	"wallpaper_server/controller"
	"wallpaper_server/utils"

	"github.com/gin-gonic/gin"
)

// token验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			controller.ReturnCommonError(c, 4004, "error", "Authorization header is required")
			c.Abort()
			return
		}
		// 解析token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := utils.ValidateToken(token)
		if err != nil {
			controller.ReturnCommonError(c, 4004, "error", "Invalid or expired token")
			c.Abort()
			return
		}
		// 将用户ID存入上下文
		c.Set("userID", userID)
		c.Next()
	}
}
