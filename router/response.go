package router

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, msg, tpl, next string, data interface{}) {
	g.C.HTML(httpCode, tpl, gin.H{
		"data": data,
		"msg":  msg,
		"next": next,
	})
	return
}
