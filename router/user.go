package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"task-test/logger"
	"task-test/model"
	"task-test/utils"
	"time"
)

var OneDayOfHours = 60 * 60 * 24

func CreateJwt(c *gin.Context) {
	//session := sessions.Default(c)
	//
	//if session.Get("")
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
			c.JSON(result.Code, gin.H{
				"result": result,
			})
		}
	} else {
		result.Message = "login field"
		result.Code = http.StatusOK
		c.JSON(result.Code, gin.H{
			"result": result,
		})
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
	re, _ := user.QueryByID()
	file, e := c.FormFile("avatar")
	if e != nil {
		//c.HTML(http.StatusOK, "error.tmpl", gin.H{
		//	"error": e,
		//})
		logger.Error("file upload field", e.Error())
		user.Avatar = re.Avatar
	} else {
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

		avatarUrl := "http://127.0.0.1:8080" + path + fileName
		user.Avatar = avatarUrl
	}
	e = user.Update()

	if e != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("can't update", e.Error())
	}
	u, _ := user.QueryByEmail()
	//fmt.Println(u.Avatar.String)
	//fmt.Println(addr)
	c.HTML(http.StatusOK, "userprofile.tmpl", gin.H{
		"user": u,
	})

}

func userLogout(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
