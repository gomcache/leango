package httpServer

import (
	"LearnGoProject/model"
	"LearnGoProject/service"
	"github.com/gin-gonic/gin"
	"time"
)

func Index(ctx *gin.Context)  {
	u := &model.User{}
	u.Username = "kinwe"
	u.Password = "hx19870527"
	u.LastLoginTime = time.Now()
	service.CreateUser(u)
	ctx.JSON(200,gin.H{"hello":"2000"})
}

