package query

import (
	"fmt"
	"github.com/vuho-pg/toolkit/check"
)

type Q map[string]interface{}

func New() Q {
	return make(Q)
}

func (q Q) Equal(name string, data interface{}) Q {
	if !check.IsZero(data) {
		q[fmt.Sprintf("%v = ?", name)] = data
	}
	return q
}

func (q Q) StartWith(name string, data string) Q {
	if !check.IsZero(data) {
		q[fmt.Sprintf("%v LIKE ?", name)] = fmt.Sprintf("%v%%", data)
	}
	return q
}

func (q Q) EndWith(name string, data string) Q {
	if !check.IsZero(data) {
		q[fmt.Sprintf("%v LIKE ?", name)] = fmt.Sprintf("%%%v", data)
	}
	return q
}

func (q Q) Contain(name string, data string) Q {
	if !check.IsZero(data) {
		q[fmt.Sprintf("%v LIKE ?", name)] = fmt.Sprintf("%%%v%%", data)
	}
	return q
}

func (q Q) In(name string, value interface{}) Q {
	if !check.IsZero(value) {
		q[fmt.Sprintf("%v IN (?)", name)] = value
	}
	return q
}
