package controller

import (
	"net/http"
	"wallpaper_server/models"

	"github.com/gin-gonic/gin"
)

type WallpaperUserController struct{}

func (wuc WallpaperUserController) Login(c *gin.Context) {
	// 验证用户名密码
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		ReturnLoginError(c, 4004, "Invalid input")
		return
	}
	token, err := models.GetUserToken(input.Username, input.Password)
	if err != nil {
		ReturnLoginError(c, 4004, token)
		return
	}
	ReturnLoginSuccess(c, 0, "success", token)
}

// 用户信息接口
func (u WallpaperUserController) Profile(c *gin.Context) {
	userID, _ := c.Get("userID")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to your profile", "userID": userID})
}
