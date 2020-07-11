package middleware

import (
	"chess2.0/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Token()gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenstring :=c.PostForm("token")

		claims := model.UserToken{}
		token,err := jwt.ParseWithClaims(tokenstring,&claims, func(token *jwt.Token) (i interface{}, err error) {
			return []byte("123"),nil
		})
		if err == nil && token.Valid{
			 c.Next()
		}else {
			 c.JSON(300,"验证失败！请先登录！")
		}

	}
}