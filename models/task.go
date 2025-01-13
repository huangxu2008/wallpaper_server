package models

import "wallpaper_server/dao"

type WallpaperCreateTask struct {
	TaskID      uint   `json:"task_id"`
	TaskName    string `json:"task_name"`
	UserID      uint   `json:"user_id"`
	ImgNum      uint   `json:"img_num"`
	ImgFormat   string `json:"img_format"`
	ImgSource   string `json:"img_source"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Parameters  string `json:"parameters"`
	CreateTime  int64  `json:"create_time"`
	UpdateTime  int64  `json:"update_time"`
}

func (WallpaperCreateTask) TableName() string {
	return "task"
}

func SetCreateTask(wct WallpaperCreateTask) (string, error) {
	result := dao.Db.Create(&wct)
	if result.Error != nil {
		return "Failed to set create task", result.Error
	}
	return "Success to set create task", nil
}
