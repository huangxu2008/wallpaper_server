package controller

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type WallpaperTaskController struct{}

func (wtc WallpaperTaskController) CreateTask(c *gin.Context) {
	// 验证用户信息
	userIdStr := c.DefaultPostForm("userId", "")
	if userIdStr == "" {
		ReturnJsonCommonError(c, 401, "usetId is empty", "error")
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		ReturnJsonCommonError(c, 401, "userId is an error param", "error")
		return
	}
	tokenId, _ := c.Get("userID")
	switch v := tokenId.(type) {
	case uint:
		if int64(v) != userId {
			ReturnJsonCommonError(c, 401, "user information mismatch", "error")
			return
		}
	default:
		ReturnJsonCommonError(c, 401, "can not get tokenId", "error")
		return
	}

	// 解析获取文件的方式
	imgFormat := c.DefaultPostForm("imgFormat", "")
	if imgFormat != "url" && imgFormat != "file" {
		ReturnJsonCommonError(c, 401, "no file upload method specified", "error")
		return
	}

	// 获取并验证必填表单数据
	widthStr := c.DefaultPostForm("width", "")
	heightStr := c.DefaultPostForm("height", "")
	imgNumStr := c.DefaultPostForm("imgNum", "")
	taskName := c.DefaultPostForm("taskName", "")
	prePromot := c.DefaultPostForm("prePromot", "")
	sufPromot := c.DefaultPostForm("sufPromot", "")
	modelName := c.DefaultPostForm("modelName", "")
	negativePrompt := c.DefaultPostForm("negativePrompt", "")
	if widthStr == "" || heightStr == "" || imgNumStr == "" || taskName == "" ||
		prePromot == "" || sufPromot == "" || modelName == "" || negativePrompt == "" {
		ReturnJsonCommonError(c, 401, "required parameter missing", "error")
		return
	}
	width, err := strconv.ParseInt(widthStr, 10, 64)
	if err != nil {
		ReturnCommonError(c, 401, "with param error", "error")
		return
	}
	height, err := strconv.ParseInt(heightStr, 10, 64)
	if err != nil {
		ReturnCommonError(c, 401, "height param error", "error")
		return
	}
	imgNum, err := strconv.ParseInt(imgNumStr, 10, 64)
	if err != nil {
		ReturnCommonError(c, 401, "imgNum param error", "error")
		return
	}

	// 获取并解析选填表单数据
	highPriorityStr := c.DefaultPostForm("highPriority", "")
	description := c.DefaultPostForm("description", "")
	imgSource := c.DefaultPostForm("imgSource", "")
	highPriority, err := strconv.ParseInt(highPriorityStr, 10, 64)
	if highPriorityStr != "" {
		if err != nil {
			ReturnCommonError(c, 401, "highPriority param error", "Params error")
			return
		}
	} else {
		highPriority = 0
	}

	// 将数据保存在数据库中，并生成一个taskId
	fmt.Printf("option params %s %s %d %d %d %d\n", description, imgSource, highPriority, width, height, imgNum)
	taskId := 100
	data := &JsonCreateTask{TaskID: uint(taskId)}
	ReturnJsonCommonSuccess(c, 200, "", "success", data)

	// 根据imgFormat决定下载文件方法
	if imgFormat == "url" {
		// 获取表单中所有名为 "urls" 的字段的值（返回一个字符串切片）
		urls := c.PostFormArray("urls")
		if len(urls) == 0 {
			ReturnCommonError(c, 401, "no urls provided", "error")
			return
		}

		// 创建文件保存目录（如果不存在）
		dir := "./uploaded_images"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			ReturnCommonError(c, 401, "unable to create directory", "error")
			return
		}

		// 创建一个无缓冲的通道来接收下载通知
		notifications := make(chan downloadImageNotification)

		// 创建一个WaitGroup来等待所有下载完成
		var wg sync.WaitGroup

		// 为每个 URL 启动一个 goroutine 来下载图片
		for _, url := range urls {
			wg.Add(1) // 增加 WaitGroup 的计数器
			go func(url string) {
				defer wg.Done()
				downloadImage(url, dir, notifications)
			}(url)
		}

		// 创建一个 goroutine 来关闭通知通道，当所有下载都完成时
		go func() {
			wg.Wait() // 等待所有下载都完成
			close(notifications)
		}()

		// 在另一个 goroutine 中监听通知通道，并处理通知
		go func() {
			for notification := range notifications {
				if notification.Error != nil {
					fmt.Printf("下载失败: URL=%s, 文件名=%s, 错误=%v\n", notification.URL, notification.FileName, notification.Error)
				} else {
					fmt.Printf("下载成功: URL=%s, 文件名=%s\n", notification.URL, notification.FileName)
				}
			}
		}()
		fmt.Printf("返回结果\n")
	} else {
		form, err := c.MultipartForm()
		if err != nil {
			ReturnCommonError(c, 4004, "error", "Unable to parse form")
			return
		}

		files := form.File["files"]
		if len(files) == 0 {
			ReturnCommonError(c, 4004, "error", "Files is empty")
			return
		}

		// 创建文件保存目录（如果不存在）
		dir := "./uploaded_images"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			ReturnCommonError(c, 4004, "error", "Unable to create directory")
			return
		}

		// 保存每个文件
		for _, file := range files {
			// 创建文件保存路径
			fileName := fmt.Sprintf("./uploaded_images/%s_%s", time.Now().Format("20060102_150405"), file.Filename)
			outFile, err := os.Create(fileName)
			if err != nil {
				ReturnCommonError(c, 4004, "error", "Unable to create file")
				return
			}
			defer outFile.Close()

			// 将文件复制到本地
			src, err := file.Open()
			if err != nil {
				ReturnCommonError(c, 4004, "error", "Unable to open file")
				return
			}
			defer src.Close()

			_, err = io.Copy(outFile, src)
			if err != nil {
				ReturnCommonError(c, 4004, "error", "Error saving file")
				return
			}
		}
	}
}
