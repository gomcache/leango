package db

import (
	"LearnGoProject/conf"
	"LearnGoProject/model"
	"LearnGoProject/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)
var db *gorm.DB
var once sync.Once

func newCoonect() *gorm.DB {

	conf := conf.Settings()
	dbURL := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",conf.Mysql.User,conf.Mysql.Password,conf.Mysql.Host,conf.Mysql.Port,conf.Mysql.DbName)
	db ,err := gorm.Open("mysql",dbURL)
	if err != nil {
		utils.System.Println(err)
	}
	db.LogMode(true)
	if conf.Mysql.ConnetMax != 0 {
		db.DB().SetMaxOpenConns(conf.Mysql.ConnetMax)
	}
	db.SingularTable(true)

	db.AutoMigrate(&model.User{})
	return db
}

func Instance() *gorm.DB  {
	once.Do(
		func() {
			db = newCoonect()
		})
	return db
}