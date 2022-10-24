package pagination

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"regexp"
	"testing"
)

func mockPagination() *Pagination {
	return &Pagination{
		Page:        rand.Int(),
		TotalPage:   0,
		TotalRecord: 0,
		Limit:       rand.Int(),
		Sort:        "+A,-B",
	}
}

func TestPagination_Gorm(t *testing.T) {
	p := mockPagination()
	sql, mock, err := sqlmock.New()
	assert.Nil(t, err)
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sql, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.Nil(t, err)
	res := make(map[string]interface{})
	mock.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf("SELECT * FROM `hello` ORDER BY A ASC,B DESC LIMIT %v OFFSET %v", p.Limit, p.Offset()))).WillReturnRows(sqlmock.NewRows([]string{"A", "B"}).AddRow(1, 2))
	db = p.Gorm(db.Table("hello")).Find(res)
	assert.Nil(t, db.Error)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestPagination_Gorm_NoLimit(t *testing.T) {
	p := mockPagination()
	sql, mock, err := sqlmock.New()
	assert.Nil(t, err)
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sql, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.Nil(t, err)
	res := make(map[string]interface{})
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `hello` ORDER BY A ASC,B DESC")).WillReturnRows(sqlmock.NewRows([]string{"A", "B"}).AddRow(1, 2))
	db = p.Gorm(db.Table("hello")).Find(res)
	assert.Nil(t, db.Error)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestPagination_Gorm_NoOrder(t *testing.T) {
	p := mockPagination()
	p.NoOrder(true)
	sql, mock, err := sqlmock.New()
	assert.Nil(t, err)
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sql, SkipInitializeWithVersion: true}), &gorm.Config{})
	assert.Nil(t, err)
	res := make(map[string]interface{})
	mock.ExpectQuery(regexp.QuoteMeta(fmt.Sprintf("SELECT * FROM `hello` LIMIT %v OFFSET %v", p.Limit, p.Offset()))).WillReturnRows(sqlmock.NewRows([]string{"A", "B"}).AddRow(1, 2))
	db = p.Gorm(db.Table("hello")).Find(res)
	assert.Nil(t, db.Error)
	assert.Nil(t, mock.ExpectationsWereMet())
}
