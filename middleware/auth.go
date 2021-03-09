package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"task-test/config"
	e "task-test/config"
	"task-test/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.Query("token")
		if token == "" {
			code = config.INVALID_PARAMS
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

//func Auth() gin.HandlerFunc {

//	return func(context *gin.Context) {
//
//		result := model.Result{
//			Code:    http.StatusUnauthorized,
//			Message: "Cannot Authorized",
//			Data:    nil,
//		}
//		//to := context.Query("token")
//
//		//fmt.Println(to)
//		auth := context.PostForm("token")
//
//		if len(auth) == 0 {
//			context.Abort()
//			context.HTML(result.Code, "error.tmpl", gin.H{
//				"error": result.Message,
//			})
//		}
//		email := context.PostForm("email")
//		//auth = strings.Fields(auth)[1]
//		// verify token
//		e := cache.Rdb.Get(context, email)
//		if e.String() == "" {
//			context.HTML(result.Code, "error.tmpl", gin.H{
//				"error": result.Message,
//			})
//		}
//
//		_, err := parseToken(auth)
//		if err != nil {
//			context.Abort()
//			result.Message = "token expired " + err.Error()
//			context.HTML(result.Code, "error.html", gin.H{
//				"error": result.Message,
//			})
//		} else {
//			logger.Info("token valid")
//		}
//		context.Next()
//	}
//}
//func parseToken(token string) (*jwt.StandardClaims, error) {
//
//	//split
//	payload := strings.Split(token, ".")
//	bytes, e := jwt.DecodeSegment(payload[1])
//
//	if e != nil {
//		logger.Error(e.Error())
//	}
//	content := ""
//	for i := 0; i < len(bytes); i++ {
//		content += string(bytes[i])
//	}
//	split := strings.Split(content, ",")
//	id := strings.SplitAfter(split[2], ":")
//	i := strings.Split(id[1], "\\u")
//	i = strings.Split(i[1], "\"")
//
//	ID, err := strconv.Atoi(i[0])
//	if err != nil {
//		logger.Error(err.Error())
//	}
//
//	user := model.User{}
//	user.Id = uint(ID)
//	u, err := user.QueryByID()
//	if err != nil {
//		logger.Error(err.Error())
//	}
//	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
//		func(token *jwt.Token) (i interface{}, e error) {
//			return []byte(utils.Configs.Jwt.Secret + u.Password), nil
//		})
//	if err == nil && jwtToken != nil {
//		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
//			return claim, nil
//		}
//	}
//	return nil, err
//}
