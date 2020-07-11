package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"size:10;not null"`
	Password string `json:"password" form:"password" gorm:"size:16;not null"`
}

type UserToken struct {
	Username string
	jwt.StandardClaims
}

//查找用户
func FindUser(mod *User)*User{
	 res :=&User{}
	tx:=DB.Begin()
	if err:=tx.Where("username = ?",mod.Username).First(&res).Error;err!=nil{
		tx.Rollback()
	}
	tx.Commit()
	return res

}

//添加用户
func CreateUser(mod *User)error{
	tx:=DB.Begin()
	if err := tx.Create(&mod).Error;err !=nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}