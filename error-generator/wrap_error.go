package error_generator

import (
	"fmt"
	"runtime"
	"strings"
)

func WrapError(err error) error {
	if err == nil {
		return nil
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}

	currentLocation := fmt.Sprintf("%s:%d", file, line)
	errString := err.Error()

	if strings.Contains(errString, "\n") || strings.Contains(errString, ": ") {
		return fmt.Errorf("%s:\n%w", currentLocation, err)
	}

	return fmt.Errorf("%s: %w", currentLocation, err)
}
