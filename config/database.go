package config

import (
	"sample/domain/model"

	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(sample_db)/sample?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
