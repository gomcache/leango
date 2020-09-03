package service

import (
	"LearnGoProject/db"
	"LearnGoProject/model"
	"fmt"
)

func CreateUser(user *model.User)  {
	db := db.Instance()
	tx := db.Begin()
	if err := tx.Create(user).Error;err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	tx.Commit()
}

func UpdateUser(user *model.User)  {
	db := db.Instance()
	tx := db.Begin()
	if err := tx.Model(user).Where("username=?", user.Username).Update(user).Error;err != nil {
		tx.Rollback()
	}
	tx.Commit()
}