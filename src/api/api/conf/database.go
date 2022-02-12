package conf

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user1:password@tcp(db)/tagudb?parseTime=True&charset=utf8mb4&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}

	return db
}