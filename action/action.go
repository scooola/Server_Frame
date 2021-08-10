package action

import (
	"Server_Frame/component"
	"Server_Frame/framework/action_service"
	"Server_Frame/framework/errors"
	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) gin.H {
	type paramStruct struct {
		Name string `json: Name`
	}
	param := paramStruct{}
	if err := c.BindJSON(&param); err != nil {
		errors.RaiseErr("Get Param Error: " + err.Error())
	}

	ret := component.HelloWorld(param.Name)
	return ret
}

func init() {
	action_service.AcitonServer.AddFunc(10000, helloWorld)
}
