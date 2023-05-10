package models

import (
	"gorm.io/gorm"
	"time"
)

type Info struct {
	gorm.Model
	UserName    string
	ServiceName string
	Login       string
	Password    string
	Expiration  time.Time
}
