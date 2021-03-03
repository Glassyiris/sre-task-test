package model

type User struct {
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `from:"password-again" binding:"eqfield=Password"`
	Nickname      string
	Avatar        string
}
