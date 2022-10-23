package pagination

import (
	"github.com/stretchr/testify/assert"
	"github.com/vuho-pg/toolkit/test"
	"testing"
)

func TestOrder_GormString(t *testing.T) {
	cases := []test.Case{
		{"desc", Order{"A", OrderDesc}, "A DESC"},
		{"asc", Order{"A", OrderAsc}, "A ASC"},
	}
	t.Parallel()
	for _, tc := range cases {
		t.Run(t.Name(), func(t *testing.T) {
			assert.Equal(t, tc.Input.(Order).GormString(), tc.Expect)
		})
	}
}

func TestOrder_splitOrder(t *testing.T) {
	cases := []test.Case{
		{"no order", "", []Order{}},
		{"1 order", "+A", []Order{{"A", OrderAsc}}},
		{"multiple order", "+A,-B,+C", []Order{{"A", OrderAsc}, {"B", OrderDesc}, {"C", OrderAsc}}},
		{"multiple order skip one", "+A,-B,*C,+D", []Order{{"A", OrderAsc}, {"B", OrderDesc}, {"D", OrderAsc}}},
	}
	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, splitOrder(tc.Input.(string)), tc.Expect)
		})
	}
}
