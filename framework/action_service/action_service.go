package action_service

import (
	"Server_Frame/framework/errors"
	"Server_Frame/share/utils"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type ActionFunc func(c *gin.Context) gin.H
type ActionFuncMap map[int]ActionFunc

type ActionService struct {
	ActionToFunc ActionFuncMap
}

var actionFunMap = make(ActionFuncMap)

var AcitonServer = &ActionService{
	ActionToFunc: actionFunMap,
}

func (act *ActionService) AddFunc(actionId int, actionFunc ActionFunc) {
	if _, ok := act.ActionToFunc[actionId]; ok {
		panic("ActionId Repeat!")
	}
	act.ActionToFunc[actionId] = actionFunc
}

func GetFunc(actionId int) ActionFunc {
	if _, ok := AcitonServer.ActionToFunc[actionId]; !ok {
		errors.RaiseErr("Server Not Have Current ActionId")
	}
	f := AcitonServer.ActionToFunc[actionId]
	return f
}

func GetActionId(c *gin.Context) int {
	var m map[string]interface{}
	b, err := c.GetRawData()
	if err != nil {
		errors.RaiseErr("GetActionId GetRawData Error: " + err.Error())
	}
	//copy request body data
	//use c.GetRawData will delete c.Request.Body own data
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	err = json.Unmarshal(b, &m)
	if err != nil {
		errors.RaiseErr("GetActionId Unmarshal Error: " + err.Error())
	}

	if _, ok := m["ActionId"]; !ok {
		errors.RaiseErr("ActionId Param Not Exist")
	}

	actIdFloat, ok := m["ActionId"].(float64)
	if !ok {
		errors.RaiseErr("Interface Type Assert Error")
	}
	actionId := int(actIdFloat)
	return actionId
}

//Response header add TraceId
func AddTraceId(c *gin.Context) string {
	//var m map[string]interface{}
	trace_id := c.GetHeader("TraceId")
	if trace_id == "" {
		trace_id = utils.CreateUUID()
	}
	c.Header("TraceId", trace_id)
	return trace_id
}

//Get TraceId from request header
//If get nil TraceId, create new TraceId and add response header
func GetTraceId(c *gin.Context) string {
	trace_id := c.GetHeader("TraceId")
	if trace_id == "" {
		trace_id = AddTraceId(c)
	}
	return trace_id
}
