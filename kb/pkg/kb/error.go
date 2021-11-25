package kb

import (
	"errors"
	"fmt"
)

func InvalidTypeError(t, f string) error {
	return errors.New(fmt.Sprintf("field \"%s\" in %s has invalid type", f, t))
}

func UnexpectedFieldError(t, f string) error {
	return errors.New(fmt.Sprintf("unexpected field \"%s\" encountered in %s", f, t))
}

func MissingRequiredFieldError(t, f string) error {
	return errors.New(fmt.Sprintf("missing required field \"%s\" in %s", f, t))
}
