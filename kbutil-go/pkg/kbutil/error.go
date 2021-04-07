package kbutil

import (
	"errors"
	"fmt"
)

func errInvalidType(t, f string) error {
	return errors.New(fmt.Sprintf("Field \"%s\" in %s has invalid type", f, t))
}

func errUnexpectedField(t, f string) error {
	return errors.New(fmt.Sprintf("Unexpected field \"%s\" encountered in %s", f, t))
}

func errMissingRequired(t, f string) error {
	return errors.New(fmt.Sprintf("Missing required field \"%s\" in %s", f, t))
}
