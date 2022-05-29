package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "user1:password@tcp(db)/tagudb?parseTime=True&charset=utf8mb4&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
