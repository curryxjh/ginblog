package main

import (
	"ginbolg/model"
	"ginbolg/routes"
)

func main() {
	// 引用数据库
	model.InitDb()

	routes.InitRouter()
}
