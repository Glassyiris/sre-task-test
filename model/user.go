package model

type User struct {
	Id       uint32 `form:"id"`
	Email    string `form:"email" binding:"email"`
	Password string `form:"password"`
	Nickname string `form:"nickname"`
	Avatar   string `form:"avatar"`
}
