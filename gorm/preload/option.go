package preload

import "gorm.io/gorm"

func Apply(ops ...Option) Options {
	return ops
}

func Normal(q string, condition ...interface{}) Option {
	return Option{
		t:         normal,
		query:     q,
		condition: condition,
	}
}

func Join(q string, condition ...interface{}) Option {
	return Option{
		t:         join,
		query:     q,
		condition: condition,
	}
}

func Custom(f func(db *gorm.DB) *gorm.DB) Option {
	return Option{
		t:      custom,
		custom: f,
	}
}
