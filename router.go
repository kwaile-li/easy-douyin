package main

import (
	"douyin/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")
	// 基础接口
	apiRouter.POST("/user/register/", controller.Register)
}
