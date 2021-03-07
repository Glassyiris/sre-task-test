package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-test/logger"
	"task-test/model"
	"task-test/utils"
	"time"
)

var OneDayOfHours = 60 * 60 * 24

func CreateJwt(c *gin.Context) {
	session := sessions.Default(c)

	user := &model.User{}
	result := &model.Result{
		Code:    200,
		Message: "login access",
		Data:    nil,
	}
	if e := c.Bind(&user); e != nil {
		result.Message = e.Error()
		result.Code = http.StatusUnauthorized
		c.JSON(http.StatusUnauthorized, gin.H{
			"result": result,
		})
	}
	u, _ := user.QueryByEmail()
	//println("user id =>", u.Id)
	if u.Password == user.Password {
		expiresTime := time.Now().Unix() + int64(OneDayOfHours)
		claims := jwt.StandardClaims{
			Audience:  user.Email,
			ExpiresAt: expiresTime,
			Id:        string(u.Id),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "iris",
			NotBefore: time.Now().Unix(),
			Subject:   "login",
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		var jwtSecret = []byte(utils.Configs.Jwt.Secret + u.Password)
		token, err := tokenClaims.SignedString(jwtSecret)
		session.Set("user", u.Email)

		err = session.Save()
		if err != nil {
			logger.Error(err.Error())
		}

		if err == nil {
			result.Message = "login access"
			result.Data = "Bearer " + token
			result.Code = http.StatusOK
			c.HTML(result.Code, "userprofile.tmpl", gin.H{
				"user": u,
			})
		} else {
			result.Message = "login field"
			result.Code = http.StatusOK
			c.HTML(result.Code, "error.tmpl", gin.H{
				"error": result.Message,
			})
		}
	} else {
		result.Message = "login field"
		result.Code = http.StatusOK
		c.HTML(result.Code, "error.tmpl", gin.H{
			"error": result.Message,
		})
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
		c.HTML(result.Code, "error.tmpl", gin.H{
			"error": result.Message,
		})
	}

	passwordAgain := c.PostForm("passwordAgain")
	if passwordAgain != user.Password {
		result.Code = http.StatusBadRequest
		result.Message = "The two passwords entered are inconsistent"
		c.HTML(result.Code, "error.tmpl", gin.H{
			"error": result.Message,
		})
	}

	_, err := user.Save()

	if err != nil {
		result.Code = http.StatusBadRequest
		result.Message = "This mailbox has already been registered"
		c.HTML(result.Code, "error.tmpl", gin.H{
			"error": result.Message,
		})
	}

	c.Redirect(http.StatusMovedPermanently, "/login")
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
		e = c.SaveUploadedFile(file, "."+path+fileName+".jpg")
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
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	u, _ := user.QueryByEmail()

	c.HTML(http.StatusOK, "userprofile.tmpl", gin.H{
		"user": u,
	})

}

func userLogout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
