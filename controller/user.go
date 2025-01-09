package controller

import (
	"wallpaper_server/models"

	"github.com/gin-gonic/gin"
)

type WallpaperUserController struct{}

func (wuc WallpaperUserController) Login(c *gin.Context) {
	// 验证用户名密码
	userName := c.DefaultPostForm("userName", "")
	password := c.DefaultPostForm("password", "")

	if userName == "" {
		ReturnJsonCommonError(c, 401, "Username is empty", "error")
		return
	}

	if password == "" {
		ReturnJsonCommonError(c, 401, "Password is empty", "error")
		return
	}

	users, token, err := models.GetUserToken(userName, password)
	if err != nil {
		ReturnJsonCommonError(c, 401, token, "error")
		return
	}

	data := &JsonLogin{NameEn: users.NameEn, NameCn: users.NameCn, UserId: users.UserID, Token: token}
	ReturnJsonCommonSuccess(c, 200, "", "success", data)
}
