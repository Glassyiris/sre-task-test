package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "just test",
	})
}

func indexLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func indexRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{})
}
