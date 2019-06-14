package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetAdmin() *gorm.DB {
	connectString := makeConnectString("admin_tool")
	return Connect(connectString)
}

func GetGame() *gorm.DB {
	connectString := makeConnectString("game")
	return Connect(connectString)
}

func Getlog() *gorm.DB {
	connectString := makeConnectString("game_log")
	return Connect(connectString)
}

func Connect(connectString string) *gorm.DB {
	db, err := gorm.Open("mysql", connectString)
	if nil != err {
		panic(err)
		return nil
	}
	return db
}
