package main

import (
	"wallpaper_server/router"
)

func main() {
	r := router.Router()

	// defer recover panic nil

	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)

	// panic("11")

	r.Run(":9999")
}
