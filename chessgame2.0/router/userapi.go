package router

import (
	"chess2.0/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//登录
func Login(c *gin.Context){

	var mod model.User
	c.Bind(&mod)

	res:=model.FindUser(&mod)
	if res.Password !=mod.Password {
		c.JSON(400,gin.H{
			"code":"400",
			"message":"密码错误或用户名错误",
		})
		return
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodES256,jwt.MapClaims{"username":mod.Username})
	tokenstring,_ := token.SignedString([]byte("123"))
	c.JSON(200,gin.H{
		"code":"200",
		"message":"登录成功",
		"data":tokenstring,
	})
	return
}

//注册
func Signup(c *gin.Context){
	var mod model.User
	c.Bind(&mod)
	result :=model.FindUser(&mod)
	if result.Username == mod.Username  {
		c.JSON(403,gin.H{
			"code":"403",
			"message":"用户名重复",
		})
		return
	}
	res:=model.CreateUser(&mod)
	if res != nil {
		c.JSON(500,gin.H{
			"code":"500",
			"message": "服务器无法处理请求",
			"data": res,
		})
		return
	}
	c.JSON(200,gin.H{
		"code":"200",
		"message": "注册成功",
	})
	return
}