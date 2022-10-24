package pagination

import (
	"fmt"
	"strings"
)

type OrderType string

const (
	OrderAsc  OrderType = "ASC"
	OrderDesc OrderType = "DESC"
)

type Order struct {
	Field string
	Order OrderType
}

func (o Order) GormString() string {
	return fmt.Sprintf("%v %v", o.Field, o.Order)
}

func splitOrder(sortString string) []Order {
	sorts := strings.Split(sortString, ",")
	orders := make([]Order, 0, len(sorts))
	for _, sort := range sorts {
		if len(sort) < 2 {
			continue
		}
		order := sort[0]
		field := sort[1:]
		switch order {
		case '-':
			orders = append(orders, Order{
				Field: field,
				Order: OrderDesc,
			})
		case '+':
			orders = append(orders, Order{
				Field: field,
				Order: OrderAsc,
			})
		default:
			continue
		}
	}
	return orders
}
