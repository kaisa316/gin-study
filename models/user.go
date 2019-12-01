package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:111111@/gin_study?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

type User struct {
	gorm.Model
	Username string
	Password string
}

//创建表结构
func CreateTable() {
	db.AutoMigrate(&User{})
	// defer db.Close()
}

//新增记录
func AddRecord(u *User) {
	// defer db.Close()
	db.Create(&u)
}

//查询根据username
func UserInfo(userName string) (user User, err error) {
	u := User{
		Username: userName,
	}
	err = db.Where(&u).First(&user).Error //绑定user object,结果会存储在这里
	return
}
