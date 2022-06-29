package model

import (
	"gorm.io/gorm"
	"time"
)

type GnModel struct {
	ID        uint64 `gorm:"primary" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}