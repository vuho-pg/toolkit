package pagination

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPagination_NoLimit(t *testing.T) {
	var p Pagination
	assert.Equal(t, p.noLimit, false)
	p.NoLimit(true)
	assert.Equal(t, p.noLimit, true)
}

func TestPagination_NoOrder(t *testing.T) {
	var p Pagination
	assert.Equal(t, p.noOrder, false)
	p.NoOrder(true)
	assert.Equal(t, p.noOrder, true)
}

func TestPagination_correct(t *testing.T) {
	var p Pagination
	assert.Equal(t, p.Page, 0)
	assert.Equal(t, p.Limit, 0)
	p.correct()
	assert.Equal(t, p.Page, DefaultPage)
	assert.Equal(t, p.Limit, DefaultLimit)
}

func TestPagination_init(t *testing.T) {
	var p Pagination
	p.Sort = "+A,-B"
	p.init()
	assert.Equal(t, p.Page, DefaultPage)
	assert.Equal(t, p.Limit, DefaultLimit)
	assert.Equal(t, p.Orders, splitOrder("+A,-B"))
	assert.Equal(t, p.ok, true)
}

func TestPagination_SetTotalRecord(t *testing.T) {
	var p Pagination
	totalRecord := rand.Int63()
	p.init()
	p.SetTotalRecord(totalRecord)
	assert.Equal(t, p.TotalRecord, totalRecord)
	totalPage := totalRecord / int64(p.Limit)
	if totalRecord%int64(p.Limit) != 0 {
		totalPage++
	}
	assert.Equal(t, p.TotalPage, totalPage+1)
}
