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
	userID, _ := c.Get("userID")
	fmt.Printf("%d profile success\n", userID)

	// 获取表单数据
	taskName := c.DefaultPostForm("taskName", "")
	userIdStr := c.DefaultPostForm("userId", "")
	modelName := c.DefaultPostForm("modelName", "")
	widthStr := c.DefaultPostForm("width", "")
	heightStr := c.DefaultPostForm("height", "")
	highPriorityStr := c.DefaultPostForm("highPriority", "true")
	description := c.DefaultPostForm("description", "")
	prePromot := c.DefaultPostForm("prePromot", "")
	sufPromot := c.DefaultPostForm("sufPromot", "")
	imgNumStr := c.DefaultPostForm("imgNum", "")
	imgFormat := c.DefaultPostForm("imgFormat", "")

	// 校验必填项
	if taskName == "" || userIdStr == "" || modelName == "" ||
		widthStr == "" || heightStr == "" || imgNumStr == "" || imgFormat == "" {
		ReturnCommonError(c, 4004, "error", "Params error")
		return
	}

	// 校验数据合法性
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		ReturnCommonError(c, 4004, "error", "Params error")
		return
	}
	width, err := strconv.ParseInt(widthStr, 10, 64)
	if err != nil {
		ReturnCommonError(c, 4004, "error", "Params error")
		return
	}
	height, err := strconv.ParseInt(heightStr, 10, 64)
	if err != nil {
		ReturnCommonError(c, 4004, "error", "Params error")
		return
	}
	highPriority, err := strconv.ParseBool(highPriorityStr)
	if err != nil {
		ReturnCommonError(c, 4004, "error", "Params error")
		return
	}
	imgNum, err := strconv.ParseInt(imgNumStr, 10, 64)
	if err != nil {
		ReturnCommonError(c, 4004, "error", "Params error")
		return
	}
	if imgFormat != "url" && imgFormat != "file" {
		ReturnCommonError(c, 4004, "error", "Params error")
		return
	}

	fmt.Printf("userId is %d\n", userId)
	fmt.Printf("width is %d\n", width)
	fmt.Printf("height is %d\n", height)
	fmt.Printf("highPriority is %t\n", highPriority)
	fmt.Printf("description is %s\n", description)
	fmt.Printf("prePromot is %s\n", prePromot)
	fmt.Printf("sufPromot is %s\n", sufPromot)
	fmt.Printf("imgNum is %d\n", imgNum)

	// 判断格式，如果上传的参数有问题，那么返回错误
	if imgFormat == "url" {

		imageURLs := []string{
			"https://iknow-pic.cdn.bcebos.com/730e0cf3d7ca7bcba60a76b1ac096b63f624a83f",
			"https://iknow-pic.cdn.bcebos.com/838ba61ea8d3fd1fa2f429a4224e251f94ca5fab",
			"https://iknow-pic.cdn.bcebos.com/d833c895d143ad4beeb4e85e90025aafa50f0691",
			"https://iknow-pic.cdn.bcebos.com/f7246b600c338744af69eee1430fd9f9d62aa0fb",
			"https://iknow-pic.cdn.bcebos.com/4afbfbedab64034f1dbe899bbdc379310a551d3f",
		}
		// 创建文件保存目录（如果不存在）
		dir := "./uploaded_images"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			ReturnCommonError(c, 4004, "error", "Unable to create directory")
			return
		}
		// 创建一个无缓冲的通道来接收下载通知
		notifications := make(chan downloadImageNotification)
		// 创建一个WaitGroup来等待所有下载完成
		var wg sync.WaitGroup
		// 为每个 URL 启动一个 goroutine 来下载图片
		for _, url := range imageURLs {
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

	ReturnCreateTaskSuccess(c, 0, "success", 1)
}
