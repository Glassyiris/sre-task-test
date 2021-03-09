package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-test/cache"
	e "task-test/config"
	"task-test/logger"
	"task-test/model"
	"task-test/utils"
)

//var ctx = context.Background()

func CreateJwt(c *gin.Context) {
	var user model.User
	resp := Gin{
		c,
	}
	err := c.BindJSON(&user)
	if err != nil {
		resp.Response(http.StatusBadRequest, e.ERROR_AUTH, "input error", "")
	}
	u, err := user.QueryByEmail()
	if err != nil {
		resp.Response(http.StatusBadRequest, e.ERROR_AUTH, "Can't found email", "")
	} else {
		token, err := utils.GenerateToken(u.Email, u.Password)
		if err != nil {
			resp.Response(http.StatusBadRequest, e.ERROR_AUTH, "Email or password wrong", "")
		}

		cache.Set(user.Email, token, 1000)
		resp.Response(http.StatusOK, e.SUCCESS, "Login Success", token)

	}
}

func userRegister(c *gin.Context) {
	var user model.User
	result := &model.Result{
		Code:    200,
		Message: "login access",
		Data:    nil,
	}
	if err := c.ShouldBind(&user); err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "You input maybe have some mistakes"
		c.HTML(result.Code, "register.tmpl", gin.H{
			"massage": result.Message,
		})
	}

	passwordAgain := c.PostForm("passwordAgain")
	if passwordAgain != user.Password {
		result.Code = http.StatusBadRequest
		result.Message = "The two passwords entered are inconsistent"
		c.HTML(result.Code, "register.tmpl", gin.H{
			"massage": result.Message,
		})
	}

	_, err := user.Save()

	if err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "This mailbox has already been registered"
		c.HTML(result.Code, "register.tmpl", gin.H{
			"massage": result.Message,
		})
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func useProfileUpdate(c *gin.Context) {
	var user model.User
	result := &model.Result{
		Code:    200,
		Message: "login access",
		Data:    nil,
	}
	if err := c.ShouldBind(&user); err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "invalid input"
		c.HTML(result.Code, "error.tmpl", gin.H{
			"error": result.Message,
		})
		logger.Error("binding error ", err.Error())
	}
	re, _ := user.QueryByID()
	file, e := c.FormFile("avatar")
	if e != nil {
		result.Code = 304
		result.Message = "invalid input"
		c.HTML(result.Code, "error.tmpl", gin.H{
			"error": result.Message,
		})
		logger.Error("file upload field", e.Error())
		user.Avatar = re.Avatar
	} else {
		//path := utils.RootPath()
		path := "/avatar/"
		fileName := utils.GetName(file.Filename)

		e = c.SaveUploadedFile(file, "."+path+fileName)
		if e != nil {
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"error": e,
			})
			logger.Error("cannot save file", e.Error())
		}

		avatarUrl := "http://127.0.0.1:8080" + path + fileName
		user.Avatar = avatarUrl
	}
	e = user.Update()

	if e != nil {
		c.HTML(http.StatusNotExtended, "error.tmpl", gin.H{
			"error": e,
		})
	}
	u, _ := user.QueryByEmail()

	c.HTML(0, "userprofile.tmpl", gin.H{
		"user": u,
	})

}

func userLogout(c *gin.Context) {
	email := c.PostForm("email")
	cache.Delete(email)

	c.Redirect(http.StatusFound, "/")
}

func profile(c *gin.Context) {
	var user model.User
	email := c.Query("email")
	user.Email = email
	u, _ := user.QueryByEmail()
	t, err := cache.Get(email)

	result := &model.Result{
		Code:    200,
		Message: "login access",
		Data:    nil,
	}
	if err != nil {
		result.Code = 404
		result.Message = "token expire"
		c.JSON(result.Code, gin.H{
			"result": result,
		})
	} else {
		result.Code = 200
		c.HTML(result.Code, "userprofile.tmpl", gin.H{
			"user":  u,
			"token": t,
		})
	}

}

func test(c *gin.Context) {
	resp := Gin{
		c,
	}
	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Printf("email: %s -> password: %s", email, password)
	resp.Response(200, 200, "Ok", "err")
}
