package controller

import "github.com/gin-gonic/gin"

type WallpaperMenuController struct{}

func (wmc WallpaperMenuController) GetMenu(c *gin.Context) {
	userID, _ := c.Get("userID")
	if userID == "" {

	}
}
