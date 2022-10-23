package error

import (
	"fmt"
	"strings"
)

type Type string

const (
	InternalError   Type = "internal"
	BadRequestError Type = "bad request"
)

const withErrorFormat = "[%v] %v"

type Error struct {
	withErrType bool
	t           Type
	err         []error
	message     string
}

func (e *Error) WithType() error {
	e.withErrType = true
	return e
}

func WithType(e error) error {
	base, ok := e.(*Error)
	if !ok {
		return e
	}
	return base.WithType()
}

func (e *Error) joinErrMessage() string {
	if len(e.err) == 0 {
		return ""
	}
	if len(e.err) == 1 {
		return e.err[0].Error()
	}
	errMsg := make([]string, 0, len(e.err))
	for _, e := range e.err {
		errMsg = append(errMsg, e.Error())
	}
	return strings.Join(errMsg, "\n")
}

func (e *Error) Error() string {
	if e.message == "" {
		if e.withErrType {
			return fmt.Sprintf(withErrorFormat, e.t, e.joinErrMessage())
		}
		return e.joinErrMessage()
	}
	if e.withErrType {
		return fmt.Sprintf(withErrorFormat, e.t, e.message)
	}
	return e.message
}

func Multiple(t Type, errs ...error) error {
	return &Error{
		t:   t,
		err: errs,
	}
}

func New(format string, args ...interface{}) error {
	return &Error{
		t:       InternalError,
		message: fmt.Sprintf(format, args...),
	}
}

func BadRequest(format string, args ...interface{}) error {
	return &Error{
		t:       BadRequestError,
		message: fmt.Sprintf(format, args...),
	}
}

func Wrap(t Type, err ...error) error {
	return &Error{
		withErrType: false,
		t:           t,
		err:         err,
	}
}
