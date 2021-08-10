package component

import "github.com/gin-gonic/gin"

func HelloWorld(name string) gin.H {
	return gin.H{"hello": "world " + name}
}
