package model

import (
	"task-test/logger"
	"task-test/utils"
)

type User struct {
	Id        uint   `form:"id" json:"id" xml:"id" gorm:"primaryKey"`
	Email     string `form:"email" json:"email" xml:"email" binding:"email"`
	Password  string `form:"password" json:"password" xml:"password"`
	Nickname  string `form:"nickname" json:"nickname" xml:"nickname"`
	Avatar    string
	CreatedAt int
	UpdatedAt int
}

func (u *User) Save() (uint, error) {
	result := utils.Db.AutoMigrate(u)
	if result != nil {
		logger.Error(result.Error())
	}
	r := utils.Db.Save(u)
	if r.Error != nil {
		logger.Error(r.Error.Error())
	}
	return u.Id, r.Error
}

func (u *User) QueryByEmail() (User, error) {
	var user User
	row := utils.Db.Where("email = ?", u.Email).Take(&user)
	if row.Error != nil {
		logger.Error(row.Error.Error())
	}
	return user, row.Error
}

func (u *User) QueryByID() (User, error) {
	var user User
	row := utils.Db.Where("id = ?", u.Id).Take(&user)
	if row.Error != nil {
		logger.Error(row.Error.Error())
	}
	return user, row.Error
}

func (u *User) Update() error {
	_, err := u.Save()
	return err
}
