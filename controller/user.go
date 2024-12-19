package controller

import (
	"wallpaper_server/models"

	"github.com/gin-gonic/gin"
)

type WallpaperUserController struct{}

func (wuc WallpaperUserController) Login(c *gin.Context) {
	// 验证用户名密码
	var input struct {
		Username string `json:"userName"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		ReturnLoginError(c, 4004, "error", "Invalid input")
		return
	}
	users, token, err := models.GetUserToken(input.Username, input.Password)
	if err != nil {
		ReturnLoginError(c, 4004, "error", token)
		return
	}
	ReturnLoginSuccess(c, 0, users.UserID, users.NameCn, token)
}
