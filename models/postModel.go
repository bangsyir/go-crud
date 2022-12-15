package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
