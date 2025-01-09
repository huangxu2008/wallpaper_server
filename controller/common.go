package controller

import "github.com/gin-gonic/gin"

type JsonCommonError struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

type JsonCommonSuccess struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Status    string      `json:"status"`
	Data      interface{} `json:"data"`
}

type JsonLogin struct {
	NameEn string `json:"nameEn"`
	NameCn string `json:"nameCn"`
	UserId uint   `json:"userId"`
	Token  string `json:"token"`
}

func ReturnJsonCommonError(c *gin.Context, errorCode int, message string, status string) {
	json := &JsonCommonError{ErrorCode: errorCode, Message: message, Status: status}
	c.JSON(200, json)
}

func ReturnJsonCommonSuccess(c *gin.Context, errorCode int, message string, status string, data interface{}) {
	json := &JsonCommonSuccess{ErrorCode: errorCode, Message: message, Status: status, Data: data}
	c.JSON(200, json)
}

type JsonCommonErrorStruct struct {
	Code    int         `json:"code"`
	Status  interface{} `json:"status"`
	Message interface{} `json:"message"`
}

type JsonLoginErrorStruct struct {
	Code    int         `json:"code"`
	Status  interface{} `json:"status"`
	Message interface{} `json:"message"`
}

type JsonLoginSuccessStruct struct {
	Code   int         `json:"code"`
	UserID uint        `json:"userId"`
	NameCn interface{} `json:"nameCn"`
	Token  interface{} `json:"token"`
}

type JsonCreateTaskErrorStruct struct {
	Code    int         `json:"code"`
	Status  interface{} `json:"status"`
	Message interface{} `json:"message"`
}

type JsonCreateTaskSuccessStruct struct {
	Code   int         `json:"code"`
	TaskId int         `json:"taskId"`
	Status interface{} `json:"status"`
}

func ReturnCommonError(c *gin.Context, code int, status interface{}, message interface{}) {
	json := &JsonCommonErrorStruct{Code: code, Status: status, Message: message}
	c.JSON(200, json)
}

func ReturnLoginError(c *gin.Context, code int, status interface{}, msg interface{}) {
	json := &JsonLoginErrorStruct{Code: code, Status: status, Message: msg}
	c.JSON(200, json)
}

func ReturnLoginSuccess(c *gin.Context, code int, id uint, name interface{}, token interface{}) {
	json := &JsonLoginSuccessStruct{Code: code, UserID: id, NameCn: name, Token: token}
	c.JSON(200, json)
}

func ReturnCreateTaskError(c *gin.Context, code int, status interface{}, message interface{}) {
	json := &JsonCreateTaskErrorStruct{Code: code, Status: status, Message: message}
	c.JSON(200, json)
}

func ReturnCreateTaskSuccess(c *gin.Context, code int, status interface{}, taskId int) {
	json := &JsonCreateTaskSuccessStruct{Code: code, Status: status, TaskId: taskId}
	c.JSON(200, json)
}
