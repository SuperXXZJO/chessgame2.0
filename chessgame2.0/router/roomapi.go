package router

import (
	"chess2.0/model"
	"github.com/gin-gonic/gin"
)

//新建房间
func Createroom(c *gin.Context){
	var mod model.Room
	c.Bind(&mod)
	res :=model.CreateRoom(&mod)
	if res != nil {
		c.JSON(500,gin.H{
			"code":"500",
			"message":"服务器无法处理",
			"data":res,
		})
		return
	}
	c.JSON(200,gin.H{
		"code":"200",
		"message":"创建成功",
	})
	return
}

//加入房间
func JoinRoom(c *gin.Context){
	var mod model.Room
	c.Bind(&mod)
	//不能同一个人加入
	res,err :=model.FindRoom(&mod)
	if err !=nil {
		c.JSON(500,gin.H{
			"code":"500",
			"message":"服务器无法处理",
			"data":err,

		})
	}
	if res.Roomowner==mod.Roomjoiner {
		c.JSON(403,gin.H{
			"code":"403",
			"message":"不能重复加入",
		})
	}
	err2 :=model.JoinRoom(&mod)
	if err2 != nil {
		c.JSON(500,gin.H{
			"code":"500",
			"message":"服务器无法处理",
			"data":err2,
		})
	}
	c.JSON(200,gin.H{
		"code":"200",
		"message":"加入成功",
	})
}

//退出房间
