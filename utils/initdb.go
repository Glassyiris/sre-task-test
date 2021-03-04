package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB
var ConfigF *Config

func initDB() {
	ConfigF = ParseConfig("./config/config.yaml")
	var err error
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", ConfigF.Database.username, ConfigF.Database.password, ConfigF.Database.ip, ConfigF.Database.port, ConfigF.Database.dataBaseName)
	Db, err = gorm.Open(mysql.Open(ConfigF.Database.Dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

}
func CreateTable() {
	initDB()
	sql := `CREATE TABLE IF NOT EXISTS users (
id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
email VARCHAR(64) NOT NULL COMMENT 'unique id',
nickname VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'user nickname, can be empty',
password VARCHAR(32) NOT NULL COMMENT 'md5 result of real password and key',
avatar VARCHAR(128)  NULL DEFAULT '' COMMENT 'user avatar, can be null',
uptime int(64) NOT NULL DEFAULT 0 COMMENT 'update time: unix timestamp',
PRIMARY KEY(id),
UNIQUE KEY email_unique (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user info table';`
	Db.Exec(sql)
}
