package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func connDB() {
	db, err = gorm.Open("mysql", "root:111111@/gin_study?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
}

type User struct {
	gorm.Model
	Username string
	Password string
}

//创建表结构
func CreateTable() {
	connDB()
	db.AutoMigrate(&User{})
	defer db.Close()
}

//新增记录
func AddRecord(u *User) {
	connDB()
	defer db.Close()
	db.Create(&u)
}

//查询根据username
func InfoByUsername() (r *gorm.DB) {
	connDB()
	defer db.Close()
	u := User{
		Username: "werwer@qq.com",
	}
	r = db.Where(&u).First(&User{})
	return
}
