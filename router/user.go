package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-test/cache"
	"task-test/logger"
	"task-test/model"
	"task-test/utils"
	"time"
)

//var ctx = context.Background()

var OneDayOfHours = 60 * 60 * 24

var token string

func CreateJwt(c *gin.Context) {
	var err error
	user := &model.User{}
	result := &model.Result{
		Code:    200,
		Message: "login access",
		Data:    nil,
	}
	if e := c.Bind(&user); e != nil {
		result.Message = e.Error()
		result.Code = http.StatusUnauthorized
		c.HTML(result.Code, "index.tmpl", gin.H{
			"error": result.Message,
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
		token, err = tokenClaims.SignedString(jwtSecret)
		cache.Set(u.Email, token, 10)
		c.SetCookie(u.Email, token, 3600, "/", "http://127.0.0.1", false, true)
		if err != nil {
			logger.Error(err.Error())
		}

		if err == nil {
			result.Message = "login access"
			result.Data = u.Email + " " + token
			result.Code = http.StatusOK
			c.HTML(result.Code, "userprofile.tmpl", gin.H{
				"token": token,
				"user":  u,
			})
		} else {
			result.Message = "login field"
			result.Code = http.StatusOK
			c.HTML(result.Code, "index.tmpl", gin.H{
				"massage": result.Message,
			})
		}
	} else {
		result.Message = "login field"
		result.Code = http.StatusOK
		c.HTML(result.Code, "index.tmpl", gin.H{
			"massage": result.Message,
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
		"token": token,
		"user":  u,
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
