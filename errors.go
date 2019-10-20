package errors

import (
	"fmt"
	"runtime"
)

func New(message string) error {
	return fmt.Errorf("%s on (%s)", message, filePath())
}

func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format+" on (%s)", append(a, filePath())...)
}

func WithPrevious(previous error, message string) error {
	return fmt.Errorf("%s on (%s) due to %w", message, filePath(), previous)
}

func filePath() string {
	pc, _, _, ok := runtime.Caller(2)
	fn := `unknown`
	if ok {
		fn = runtime.FuncForPC(pc).Name()
	}

	return fmt.Sprintf(`%s`, fn)
}
