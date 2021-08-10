//Auther: scola
//Date: 2021/08/08 20:20
//Description:
//Σ(っ °Д °;)っ

package errors

import (
	"Server_Frame/share/enum/error_code"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"time"
)

type ExtType map[string]interface{}

type Error interface {
	Error() string
	Code() int
	Ext() ExtType
}

type ServerError struct {
	msg  string
	code int
	ext  ExtType
}

func (e ServerError) Error() string {
	return e.msg
}

func (e ServerError) Code() int {
	return e.code
}

func (e ServerError) Ext() ExtType {
	if e.ext == nil {
		var ext = make(ExtType)
		return ext
	}
	return e.ext
}

func NewServerErr(code int, msg string) ServerError {
	serverEror := ServerError{
		msg:  msg,
		code: code,
	}
	return serverEror
}

func NewServerErrExt(code int, msg string, ext ExtType) ServerError {
	serverEror := ServerError{
		msg:  msg,
		code: code,
		ext:  ext,
	}
	return serverEror
}

func RaiseErr(msg string) {
	err := NewServerErr(error_code.Failed, msg)
	panic(err)
}

func RaiseErrExt(msg string, ext ExtType) {
	err := NewServerErrExt(error_code.Failed, msg, ext)
	panic(err)
}

func LogError(err interface{}) {
	timeObj := time.Now()
	timeStr := timeObj.Format("2006-01-02 03:04:05")
	red := color.New(color.FgRed).PrintfFunc()
	e, ok := err.(ServerError)
	if ok {
		//server error
		red("\n[SERVER ERROR]")
		fmt.Println(timeStr)
		fmt.Println("Err Msg:", e.Error(), "\nErr Code:", e.Code(), "\nExt Msg:", e.Ext())
	} else {
		//other error
		red("\n[OHTER ERROR]")
		fmt.Println(timeStr)
		fmt.Println("Err Msg:", err)
	}
	fmt.Println(string(debug.Stack()))
	//TODO: add error msg to log
}

// only cathch error and log error
func CatchError() {
	if err := recover(); err != nil {
		LogError(err)
	}
}

// catch,log error and return error msg with success status
func CatchRetError(c *gin.Context) {
	if err := recover(); err != nil {
		//log error
		LogError(err)

		//return error msg with success status
		e, ok := err.(ServerError)
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"Code":  e.Code(),
				"Error": e.Error(),
				"Ext":   e.Ext(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Code":     error_code.Failed,
				"Error":    "Unknow Error",
				"ErrStack": string(debug.Stack()),
			})
		}
	}
}
