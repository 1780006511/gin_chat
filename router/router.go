package router

import (
	"gin_chat/docs"
	"gin_chat/service"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// 1.创建路由
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.GET("/user/findUserByNameAndPassword", service.FindUserByNameAndPassword)
	r.POST("/user/updateUser", service.UpdateUser)
	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)
	return r
}
