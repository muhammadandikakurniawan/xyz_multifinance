package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(connString string, config gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(connString), &config)
	if err != nil {
		return nil, err
	}
	return db, nil
}
