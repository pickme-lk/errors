package errors

import (
	"errors"
	"fmt"
	"runtime"
)

type Error struct {
	previous error
	current  error
}

func (e *Error) Error() string {

	if e.previous != nil {
		//return fmt.Sprintf("%s due to\n%s", e.current.Error(),  e.previous.Error())
		return fmt.Sprintf("%s due to %s", e.current.Error(), e.previous.Error())
	}

	return e.current.Error()
}

func (e Error) Errors() []error {
	return nil
}

func New(prefix string, message interface{}) error {
	return &Error{
		current: errors.New(fmt.Sprintf(`(%s) (%+v%s)`, prefix, message, filePath())),
	}
}

func WithPrevious(previous error, prefix string, message interface{}) error {

	current := errors.New(
		fmt.Sprintf(`(%s) (%+v%s)`, prefix, message, filePath()))

	return &Error{
		previous: previous,
		current:  current,
	}
}

func filePath() string {
	_, f, l, ok := runtime.Caller(2)
	if !ok {
		f = `<Unknown>`
		l = 1
	}

	return fmt.Sprintf(` on %s:%d`, f, l)
}
