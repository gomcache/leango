package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;type:varchar(18)"`
	Password string `gorm:"type:varchar(18)"`
	Status  int   //-1: 不可用 0:新用户 1:完成注册，未审批 2:完成审批
	LastLoginTime time.Time
}
