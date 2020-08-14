package router

import (
	"github.com/gin-gonic/gin"
	"awesomeProject/tron-go/api"
)

func Router()  *gin.Engine{

	router := gin.Default()
	v1:=router.Group("v1")
	v1.GET("/GetAccount/:address", api.GetAccount)
	v1.GET("/GetAddress/:address/:num", api.GetAdress)
	v1.GET("/CreateAddress/:hexkey", api.CreateAddress)
	v1.GET("/TrcTransfer/:from/:to/:value", api.TrcTransfer)
	v1.GET("/Trc20Transfer/:from/:to/:value", api.Trc20Transfer)
	return router
}