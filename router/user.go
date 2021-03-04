package router

import (
	"fmt"
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
		logger.Error(err.Error())
	}

	passwordAgain := c.PostForm("passwordAgain")
	if passwordAgain != user.Password {
		c.String(http.StatusBadRequest, "You should put the same password twice")
		logger.Error("the twice input are different")
		return
	}

	_, err := user.Save()

	if err != nil {
		c.String(http.StatusOK, "this email already signed")
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/login")
}

func useProfileUpdate(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		logger.Error("binding error ", err.Error())
	}
	file, e := c.FormFile("avatar")
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		logger.Error("file upload field", e.Error())
	}
	//path := utils.RootPath()
	path := "/avatar/"
	fileName := file.Filename
	e = c.SaveUploadedFile(file, "."+path+fileName)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		logger.Error("cannot save file", e.Error())
	}

	avatarUrl := path + fileName
	user.Avatar = avatarUrl
	e = user.Update()

	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("can't update", e.Error())
	}
	u, _ := user.QueryByEmail()
	//fmt.Println(u.Avatar.String)
	addr := "http://127.0.0.1:8080" + u.Avatar
	fmt.Println(addr)
	c.HTML(http.StatusOK, "userprofile.tmpl", gin.H{
		"image": addr,
		"user":  u,
	})

}

func userLogout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
