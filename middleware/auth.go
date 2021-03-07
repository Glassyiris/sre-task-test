package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"task-test/logger"
	"task-test/model"
	"task-test/utils"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {

		result := model.Result{
			Code:    http.StatusUnauthorized,
			Message: "Cannot Authorized",
			Data:    nil,
		}
		auth := context.Request.Header.Get("Authorization")

		if len(auth) == 0 {
			context.Abort()
			context.HTML(result.Code, "error.tmpl", gin.H{
				"error": result.Message,
			})
		}

		auth = strings.Fields(auth)[1]
		// verify token
		_, err := parseToken(auth)
		if err != nil {
			context.Abort()
			result.Message = "token expired " + err.Error()
			context.HTML(result.Code, "error.html", gin.H{
				"error": result.Message,
			})
		} else {
			logger.Info("token valid")
		}
		context.Next()
	}
}
func parseToken(token string) (*jwt.StandardClaims, error) {

	//split
	payload := strings.Split(token, ".")
	bytes, e := jwt.DecodeSegment(payload[1])

	if e != nil {
		logger.Error(e.Error())
	}
	content := ""
	for i := 0; i < len(bytes); i++ {
		content += string(bytes[i])
	}
	split := strings.Split(content, ",")
	id := strings.SplitAfter(split[2], ":")
	i := strings.Split(id[1], "\\u")
	i = strings.Split(i[1], "\"")

	ID, err := strconv.Atoi(i[0])
	if err != nil {
		logger.Error(err.Error())
	}

	user := model.User{}
	user.Id = uint(ID)
	u, err := user.QueryByID()
	if err != nil {
		logger.Error(err.Error())
	}
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte(utils.Configs.Jwt.Secret + u.Password), nil
		})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
