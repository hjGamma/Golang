package main

import (
	"web-gin/databases"
	"web-gin/routers"
)

// 定义中间

func main() {
	router := routers.InitRouter()
	//初始化数据库
	databases.InitDatabase()

	//静态资源
	router.Static("/static", "./static")
	router.Run(":8000")
}
