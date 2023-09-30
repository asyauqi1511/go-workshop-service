package errorwrapper

import (
	"fmt"
	"runtime"
)

func Wrap(err error) error {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		return fmt.Errorf("%v %s:%d", err, filename, line)
	}
	return nil
}
