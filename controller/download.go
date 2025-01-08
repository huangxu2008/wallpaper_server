package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type downloadImageNotification struct {
	URL      string
	FileName string
	Error    error
}

func downloadImage(url string, destinationDir string, notifications chan<- downloadImageNotification) {
	// 从URL中解析出文件名
	fileName := filepath.Base(url)
	// 检查文件名是否包含后缀名
	ext := strings.TrimPrefix(filepath.Ext(fileName), ".")
	if ext == "" {
		// 如果没有后缀名，则添加.png
		fileName += ".png"
	}
	// 创建目标文件路径
	destinationFile := filepath.Join(destinationDir, fileName)
	// 发送 http get 请求
	resp, err := http.Get(url)
	if err != nil {
		// 发送包含错误的通知
		notifications <- downloadImageNotification{URL: url, FileName: fileName, Error: err}
		return
	}
	defer resp.Body.Close()
	// 检查 http 响应状态码
	if resp.StatusCode != http.StatusOK {
		// 发送包含错误的通知
		notifications <- downloadImageNotification{URL: url, FileName: fileName, Error: fmt.Errorf("Http status code: %d", resp.StatusCode)}
		return
	}
	// 创建目标文件
	outFile, err := os.Create(destinationFile)
	if err != nil {
		// 发送包含错误的通知
		notifications <- downloadImageNotification{URL: url, FileName: fileName, Error: err}
		return
	}
	defer outFile.Close()
	// 将响应体写入目标文件
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		// 发送包含错误的通知
		notifications <- downloadImageNotification{URL: url, FileName: fileName, Error: err}
		return
	}
	// 发送成功的通知
	notifications <- downloadImageNotification{URL: url, FileName: fileName, Error: nil}
}
