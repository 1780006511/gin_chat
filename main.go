package main

import (
	"gin_chat/router"
	"gin_chat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	r := router.Router()
	r.Run(":8000")
}
