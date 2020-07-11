package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func init(){
	db,err:=gorm.Open("mysql","root:root@/chess?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return
	}
	DB=db
	db.AutoMigrate(&User{},&Room{})




}

