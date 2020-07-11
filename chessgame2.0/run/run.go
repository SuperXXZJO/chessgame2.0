package main

import (
	"chess2.0/game"
	"chess2.0/router"
	"github.com/gin-gonic/gin"
)

func run(){
	r := gin.Default()
	r.POST("/signup",router.Signup)   //注册



	//v1:=r.Group("/v1",middleware.Token())
	r.POST("/login",router.Login)   //登录
	r.POST("/room",router.Createroom)//创建房间
	r.POST("/room/:roomname",router.JoinRoom)//加入房间
	r.POST("/room/:roomname/start",game.Startgame)//开始游戏
	r.Run()
}
