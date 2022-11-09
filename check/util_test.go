package check

import (
	"github.com/stretchr/testify/assert"
	"github.com/vuho-pg/toolkit/test"
	"testing"
)

func TestIsZero(t *testing.T) {
	type mockStruct struct {
		A string
		B int
	}
	cases := []test.Case{
		{"nil", nil, true},
		{"arr", []interface{}{1, "2"}, false},
		{"empty struct", mockStruct{}, true},
		{"int 0", 0, true},
		{"int", 1, false},
	}
	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, IsZero(tc.Input), tc.Expect)
		})
	}
}
