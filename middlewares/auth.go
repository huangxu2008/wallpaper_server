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
			controller.ReturnJsonCommonError(c, 401, "Authorization header is required", "error")
			c.Abort()
			return
		}
		// 解析token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := utils.ValidateToken(token)
		if err != nil {
			controller.ReturnJsonCommonError(c, 401, "Invalid or expired token", "error")
			c.Abort()
			return
		}
		// 将用户ID存入上下文
		c.Set("userID", userID)
		c.Next()
	}
}
