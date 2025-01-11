package controller

import (
	"wallpaper_server/models"

	"github.com/gin-gonic/gin"
)

type WallpaperMenuController struct{}

func (wmc WallpaperMenuController) GetMenu(c *gin.Context) {
	menuList, msg, err := models.GetMenuList()
	if err != nil {
		ReturnJsonCommonError(c, 401, msg, "error")
		return
	}

	ReturnJsonCommonSuccess(c, 200, "", "success", menuList)
}
