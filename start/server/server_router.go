//Auther: scola
//Date: 2021/08/08 20:04
//Description:
//Σ(っ °Д °;)っ

package server

import (
	"Server_Frame/framework/action_service"
	"Server_Frame/framework/errors"
	"Server_Frame/infrastructure/config"
	"Server_Frame/share/enum/error_code"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type pushServer struct {
	Router *gin.Engine
	Config config.ServerConfig
}

var PushServer pushServer

func initRouter(ginMode string) {
	gin.SetMode(ginMode)
	PushServer.Router = gin.Default()
}

func routerRun(port int) {
	portStr := ":" + strconv.Itoa(port)
	PushServer.Router.Run(portStr)
}

func initFrontend() {
	PushServer.Router.POST("/frontend", func(c *gin.Context) {
		defer errors.CatchRetError(c)
		//add response header trace_id
		action_service.AddTraceId(c)
		//get action_id
		actionId := action_service.GetActionId(c)
		//get func
		f := action_service.GetFunc(actionId)
		response := f(c)
		response["code"] = error_code.Success
		c.JSON(http.StatusOK, response)
	})
}

func initHealthCheck() {
	healchCheckFunc := func(c *gin.Context) {
		defer errors.CatchRetError(c)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"code":   error_code.Success,
		})
	}
	PushServer.Router.GET("/healch_check", healchCheckFunc)
	PushServer.Router.POST("/healch_check", healchCheckFunc)
}

func initConfig() {
	PushServer.Config = config.GetJsonConf()
}
