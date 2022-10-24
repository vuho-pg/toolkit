package pagination

import "gorm.io/gorm"

// Gorm : apply the pagination to existed gorm.DB
func (p *Pagination) Gorm(db *gorm.DB) *gorm.DB {
	if !p.ok {
		p.init()
	}
	if !p.noLimit {
		db = db.Limit(p.Limit).Offset((p.Page - 1) * p.Limit)
	}
	if !p.noOrder {
		if len(p.Orders) != 0 {
			for _, order := range p.Orders {
				db = db.Order(order.GormString())
			}
		}
	}
	return db
}
