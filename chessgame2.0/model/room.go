package model

import "github.com/jinzhu/gorm"

type Room struct {
	gorm.Model
	Roomname string   `gorm:"not null"`
	Roomowner string
	Roomjoiner  string
}

//新建一个房间
func CreateRoom(mod *Room)error{
	tx:=DB.Begin()
	if err :=tx.Create(&mod).Error;err!=nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//加入房间
func JoinRoom(mod *Room)error{
	tx :=DB.Begin()
	err :=DB.Model(mod).Where("roomname=?",mod.Roomname).Update("roomjoiner",mod.Roomjoiner).Error
	if err!=nil{
		tx.Rollback()
		return err
	}
	return nil
}

//查询房间
func FindRoom(mod *Room)(*Room,error){
	var res *Room
	tx :=DB.Begin()
	if err:=tx.Where("roomname=?",mod.Roomname).First(res).Error;err !=nil{
		tx.Rollback()
		return &Room{},err
	}
	return res,nil
}