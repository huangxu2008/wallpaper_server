package main

import (
	"wallpaper_server/router"
)

func main() {
	r := router.Router()
	if r == nil {
		/// to modify 输出日志，报错，退出服务，记录时间
	} else {
		/// to modify 输出日志，正确，启动服务，记录时间
		r.Run(":9999")
	}
}
