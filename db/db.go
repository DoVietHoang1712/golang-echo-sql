package db

import (
	"fmt"
	"golang-echo-sql/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dsn := "root:17122000@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	fmt.Println(dsn)
	// db.Create(&model.User{Username: "123", Password: "456", Email: "1111"})
	var user *model.User
	db.Where(&model.User{Username: "1"}).Attrs(&model.User{Password: "123", Email: "provip"}).FirstOrCreate(&user)
	fmt.Println(user)
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
