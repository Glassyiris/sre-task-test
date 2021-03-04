package model

import (
	"database/sql"
	"task-test/logger"
	"task-test/utils"
	"time"
)

type User struct {
	Id        uint           `form:"id" json:"id" xml:"id" gorm:"primaryKey"`
	Email     string         `form:"email" json:"email" xml:"email" binding:"email"`
	Password  string         `form:"password" json:"password" xml:"password"`
	Nickname  string         `form:"nickname" json:"nickname" xml:"nickname"`
	Avatar    sql.NullString `form:"avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Save() uint {
	result := utils.Db.AutoMigrate(u)
	if result != nil {
		logger.Error(result.Error())
	}
	r := utils.Db.Save(u)
	if r.Error != nil {
		logger.Error(r.Error.Error())
	}
	return u.Id
}

func (u *User) QueryByEmail() (User, error) {
	var user User
	row := utils.Db.Where("email = ?", u.Email).Take(&user)
	if row.Error != nil {
		logger.Error(row.Error.Error())
	}
	return user, row.Error
}

func (u *User) QueryByID(id int64) (User, error) {
	var user User
	row := utils.Db.Where("id = ?", u.Id).Take(&user)
	if row.Error != nil {
		logger.Error(row.Error.Error())
	}
	return user, row.Error
}

func (u *User) Update(id int64) error {
	result := utils.Db.Save(u)

	return result.Error
}
