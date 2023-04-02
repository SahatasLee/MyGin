package controllers

import "gorm.io/gorm"

type DBController struct {
	Database *gorm.DB
}
