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

type nilElementError struct{}

func (e *nilElementError) Error() string {
	return "element cannot be nil"
}

type invalidTagError struct {
	expected string
	actual   string
}

func (e *invalidTagError) Error() string {
	return fmt.Sprintf("invalid element tag: expected = %q, actual = %q", e.expected, e.actual)
}

type missingRequiredAttributeError struct {
	element   string
	attribute string
}

func (e *missingRequiredAttributeError) Error() string {
	return fmt.Sprintf("missing required attribute: element = %q, attribute = %q", e.element, e.attribute)
}

type invalidAttributeTypeError struct {
	element   string
	attribute string
}

func (e *invalidAttributeTypeError) Error() string {
	return fmt.Sprintf("invalid attribute type: element = %q, attribute = %q", e.element, e.attribute)
}

type unexpectedAttributeError struct {
	element   string
	attribute string
}

func (e *unexpectedAttributeError) Error() string {
	return fmt.Sprintf("unexpected attribute: element = %q, attribute = %q", e.element, e.attribute)
}

type invalidChildElementError struct {
	element string
	child   string
}

func (e *invalidChildElementError) Error() string {
	return fmt.Sprintf("invalid child: element = %q, child = %q", e.element, e.child)
}
