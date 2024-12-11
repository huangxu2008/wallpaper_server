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

	db, err := gorm.Open("mysql", config.Mysqldb)
	if err != nil {
		x := 1
		x++
	}
	db.AutoMigrate(&models.TokenUser{})
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		user.GET("/info/:id", controller.UserController{}.GetUserInfo)
		user.POST("/list", controller.UserController{}.GetList)
		user.POST("/add", controller.UserController{}.AddUser)
		user.POST("/update", controller.UserController{}.UpdateUser)
		user.POST("/delete", controller.UserController{}.DeleteUser)
		user.POST("/list/test", controller.UserController{}.GetUserListTest)
		user.POST("/login", controller.UserController{}.Login(db))
	}

	protected := r.Group("/api", middlewares.AuthMiddleware())
	{
		protected.GET("/profile", controller.UserController{}.Profile)
	}

	order := r.Group("/order")
	{
		order.POST("/list", controller.OrderController{}.GetList)
	}

	return r
}
