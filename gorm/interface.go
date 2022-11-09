package gorm

import "gorm.io/gorm"

type Applicable interface {
	Apply(db *gorm.DB) *gorm.DB
}
