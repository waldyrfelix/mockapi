package models

import (
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Actor struct {
	ActorId    int        `gorm:"column:actor_id;PRIMARY_KEY;AUTO_INCREMENT"`
	FirstName  string     `gorm:"column:first_name;type:varchar(45);not null"`
	LastName   string     `gorm:"column:last_name;type:varchar(45);not null"`
	LastUpdate *time.Time `gorm:"column:last_update;type:timestamp;not null"`
}
