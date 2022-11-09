package query

import "fmt"

type Q map[string]interface{}

func New() Q {
	return make(Q)
}

func (q *Q) Equal(name string, data interface{}) *Q {
	(*q)[fmt.Sprintf("%v = ?", name)] = data
	return q
}

func (q *Q) StartWith(name string, data string) *Q {
	(*q)[fmt.Sprintf("%v LIKE ?")] = fmt.Sprintf("%v%", data)
}
