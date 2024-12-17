package models

import (
	"time"
	"wallpaper_server/dao"
	"wallpaper_server/utils"
)

type WallpaperUser struct {
	UserID    uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Username  string     `gorm:"unique" json:"username"`
	Password  string     `json:"password"`
}

func (WallpaperUser) TableName() string {
	return "user"
}

func GetUserToken(username string, password string) (string, error) {
	var users WallpaperUser
	if err := dao.Db.Where("username = ? AND password = ?", username, password).First(&users).Error; err != nil {
		return "Invalid username or password", err
	}
	// 生成token
	token, err := utils.GenerateToken(users.UserID)
	if err != nil {
		return "Failed to generate token", err
	}
	return token, err
}
