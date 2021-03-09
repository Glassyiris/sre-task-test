package router

import "github.com/gin-gonic/gin"

func Index(c *gin.Context)  {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "User manager",
	})
}

func indexRegister(c *gin.Context) {
	c.HTML(200, "register.tmpl", gin.H{
		"title": "User manager",
	})
}
