package db

import (
	"github.com/nelsonkti/echo-framework/lib/db/mysql"
	"gorm.io/gorm"
)

func Connect(name string) *gorm.DB {
	return mysql.Session(name)
}
