package models

import (
	"wallpaper_server/dao"
	"wallpaper_server/utils"
)

type WallpaperUser struct {
	UserID     uint   `gorm:"primaryKey;autoIncrement"`
	RoleID     *int   `gorm:"default:null"`
	NameEn     string `gorm:"size:128;unique;index"`
	NameCn     string `gorm:"size:128"`
	Password   string `gorm:"size:128"`
	CreateTime int64  `gorm:"default:null"`
	UpdateTime int64  `gorm:"default:null"`
}

func (WallpaperUser) TableName() string {
	return "user"
}

func GetUserToken(username string, password string) (WallpaperUser, string, error) {
	var users WallpaperUser
	if err := dao.Db.Where("name_en = ? AND password = ?", username, password).First(&users).Error; err != nil {
		return users, "Invalid username or password", err
	}
	// 生成token
	token, err := utils.GenerateToken(users.UserID)
	if err != nil {
		return users, "Failed to generate token", err
	}
	return users, token, err
}
