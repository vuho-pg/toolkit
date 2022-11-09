package preload

import "gorm.io/gorm"

type preloadType int

const (
	normal preloadType = iota
	join
	custom
)

type Options []Option

func (o Options) Apply(db *gorm.DB) *gorm.DB {
	for _, option := range o {
		db = option.Apply(db)
	}
	return db
}

type Option struct {
	t         preloadType
	query     string
	condition []interface{}
	custom    func(db *gorm.DB) *gorm.DB
}

func (o Option) Apply(db *gorm.DB) *gorm.DB {
	switch o.t {
	case normal:
		return db.Preload(o.query, o.condition...)
	case join:
		return db.Joins(o.query, o.condition...)
	case custom:
		return db.Preload(o.query, o.custom)
	default:
		return db
	}
}
