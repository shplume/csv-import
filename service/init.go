package service

import (
	"github.com/shplume/csv-import/model"
	"gorm.io/gorm"
)

var (
	data   []model.VideoData
	dbConn *gorm.DB
)

func init() {
	dbConn = model.GetDBConnection()
}
