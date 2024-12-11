package controller

import (
	"net/http"
	"strconv"
	"wallpaper_server/models"
	"wallpaper_server/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")

	id, _ := strconv.Atoi(idStr)

	user, _ := models.GetUserTest(id)

	ReturnSuccess(c, 0, name, user, 1)
}

func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "保存错误")
		return
	}
	ReturnSuccess(c, 0, "保存成功", id, 1)
}

func (u UserController) UpdateUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.UpdateUser(id, username)
	if err != nil {
		ReturnError(c, 4002, "更新失败")
		return
	}
	ReturnSuccess(c, 0, "更新成功", true, 1)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4002, "删除失败")
		return
	}
	ReturnSuccess(c, 0, "删除成功", true, 1)
}

func (u UserController) GetList(c *gin.Context) {
	// logger.Write("日志信息", "user")

	num1 := 1
	num2 := 0
	num3 := num1 / num2

	ReturnError(c, 4004, num3)
}

func (u UserController) GetUserListTest(c *gin.Context) {
	users, err := models.GetUserListTest()
	if err != nil {
		ReturnError(c, 4004, "没有相关数据")
		return
	}
	ReturnSuccess(c, 0, "获取成功", users, 1)
}

func (u UserController) Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// 验证用户
		var users models.TokenUser
		if err := db.Where("username = ? AND password = ?", input.Username, input.Password).First(&users).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		// 生成 Token
		token, err := utils.GenerateToken(users.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// 用户信息接口
func (u UserController) Profile(c *gin.Context) {
	userID, _ := c.Get("userID")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to your profile", "userID": userID})
}
