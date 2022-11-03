package errors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	e := New("internal error: %v", 1)
	assert.Equal(t, e.Error(), fmt.Sprintf("internal error: %v", 1))
	assert.Equal(t, e.(*Error).WithType().Error(), fmt.Sprintf(withErrorFormat, InternalError, fmt.Sprintf("internal error: %v", 1)))
}

func TestBadRequest(t *testing.T) {
	e := BadRequest("bad request: %v", 1)
	assert.Equal(t, e.Error(), fmt.Sprintf("bad request: %v", 1))
	assert.Equal(t, e.(*Error).WithType().Error(), fmt.Sprintf(withErrorFormat, BadRequestError, fmt.Sprintf("bad request: %v", 1)))
}

func TestMultiple(t *testing.T) {
	err0 := Multiple(BadRequestError)
	assert.Equal(t, err0.Error(), "")
	err1 := Multiple(BadRequestError, BadRequest("1"))
	assert.Equal(t, err1.Error(), "1")
	errMulti := Multiple(BadRequestError, errors.New("1"), errors.New("2"))
	assert.Equal(t, errMulti.Error(), "1\n2")
}

func TestError_WithType(t *testing.T) {
	err1 := Wrap(BadRequestError, errors.New("1"))
	err2 := errors.New("2")
	assert.Equal(t, err1.Error(), "1")
	assert.Equal(t, WithType(err1).Error(), fmt.Sprintf(withErrorFormat, BadRequestError, "1"))
	assert.Equal(t, WithType(err2).Error(), "2")
}
