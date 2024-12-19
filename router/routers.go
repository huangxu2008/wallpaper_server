package router

import (
	"wallpaper_server/config"
	"wallpaper_server/controller"
	"wallpaper_server/middlewares"
	"wallpaper_server/models"
	"wallpaper_server/pck/logger"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Router() *gin.Engine {
	var r *gin.Engine
	db, err := gorm.Open("mysql", config.Mysqldb)
	if err == nil {
		db.AutoMigrate(&models.WallpaperUser{})
		r = gin.Default()

		// 日志配置
		r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
		r.Use(logger.Recover)

		// 设置路由
		user := r.Group("wallpaper/user")
		{
			user.POST("/login", controller.WallpaperUserController{}.Login)
		}
		task := r.Group("wallpaper/task", middlewares.AuthMiddleware())
		{
			task.POST("/taskCreate", controller.WallpaperTaskController{}.CreateTask)
		}

	}
	return r
}
