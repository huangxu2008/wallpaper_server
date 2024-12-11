package controller

import "github.com/gin-gonic/gin"

type OrderController struct{}

type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetList(c *gin.Context) {
	// 这是用参数方式请求的
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "wangwu")
	// ReturnSuccess(c, 0, cid, name, 1)

	// 这是用json方式请求的
	// param := make(map[string]interface{})
	// err := c.BindJSON(&param)
	// if err == nil {
	// 	ReturnSuccess(c, 0, "success", param, 1)
	// 	return
	// }
	// ReturnError(c, 4001, gin.H{"err": err})

	// 这是提前定义结构体
	search := &Search{}
	err := c.BindJSON(&search)
	if err == nil {
		ReturnSuccess(c, 0, search.Name, search.Cid, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err})
}
