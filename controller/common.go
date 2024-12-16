package controller

import "github.com/gin-gonic/gin"

type JsonLoginErrorStruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

type JsonLoginSuccessStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Token interface{} `json:"token"`
}

func ReturnLoginError(c *gin.Context, code int, msg interface{}) {
	json := &JsonLoginErrorStruct{Code: code, Msg: msg}
	c.JSON(200, json)
}

func ReturnLoginSuccess(c *gin.Context, code int, msg interface{}, token interface{}) {
	json := &JsonLoginSuccessStruct{Code: code, Msg: msg, Token: token}
	c.JSON(200, json)
}
