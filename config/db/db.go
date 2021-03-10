package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/waldyrfelix/mockapi/config"
)

func OpenDb() (db *gorm.DB, err error) {
	config := config.GetConfig()
	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name)

	return gorm.Open(config.DB.Dialect, uri)
}
