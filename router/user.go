package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"task-test/logger"
	"task-test/model"
	"task-test/utils"
	"time"
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
	path := utils.RootPath()
	path = path + "avatar/"
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		logger.Error("can't create folder", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	e = c.SaveUploadedFile(file, path+fileName)
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		logger.Error("cannot save file", e.Error())
	}
	avatarUrl := "http://localhost:8080/avatar/" + fileName
	user.Avatar = sql.NullString{String: avatarUrl}
	e = user.Update()
	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("can't update", e.Error())
	}
	u, _ := user.QueryByID()

	c.HTML(http.StatusOK, "userprofile.tmpl", gin.H{
		"user": u,
	})

}

func userLogout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
