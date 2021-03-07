package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"task-test/config"
	"task-test/logger"
)

var Db *gorm.DB
var Configs *config.Config

func InitDB() {
	var err error
	Configs = config.ParseConfig("./config")
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", ConfigF.Database.username, ConfigF.Database.password, ConfigF.Database.ip, ConfigF.Database.port, ConfigF.Database.dataBaseName)
	Db, err = gorm.Open(mysql.Open(Configs.Database.Dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err.Error())
	}

}

//func CreateTable() {
//	InitDB()
//	sql := `CREATE TABLE IF NOT EXISTS users (
//id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
//email VARCHAR(64) NOT NULL COMMENT 'unique id',
//nickname VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'user nickname, can be empty',
//password VARCHAR(32) NOT NULL COMMENT 'md5 result of real password and key',
//avatar VARCHAR(128)  NULL DEFAULT '' COMMENT 'user avatar, can be null',
//PRIMARY KEY(id),
//UNIQUE KEY email_unique (email)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user info table';`
//	Db.Exec(sql)
//}
//
