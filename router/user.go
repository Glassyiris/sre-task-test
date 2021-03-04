package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-test/logger"
	"task-test/model"
)

func userLogin(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		logger.Info("cannot bind")
		c.String(http.StatusOK, "Login field")
	}
	u, err := user.QueryByEmail()

	if err != nil {
		c.String(http.StatusOK, "this account not found")
		return
	}

	if u.Password == user.Password {
		c.HTML(http.StatusOK, "userprofile.tmpl", gin.H{
			"user": u,
		})
	} else {
		c.String(http.StatusOK, "Password or email has mistake")
	}
}

func userRegister(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, "input void")
		log.Panicln("err ->", err.Error())
	}
	passwordAgain := c.PostForm("passwordAgain")
	if passwordAgain != user.Password {
		c.String(http.StatusBadRequest, "You should put the same password twice")
		logger.Error("the twice input are different")
		return
	}
	user.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}

func useProfileUpdate(c *gin.Context) {

}

func userLogout(c *gin.Context) {

}
