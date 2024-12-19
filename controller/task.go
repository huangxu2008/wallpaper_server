package controller

import (
	"fmt"
	"io"
	"os"
	"strconv"
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
