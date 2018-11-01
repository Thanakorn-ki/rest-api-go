package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(db:3306)/seller?charset=utf8")

	if err != nil {
		panic(err)
	}
	return db
}
