package unmarshal

import (
	"fmt"
)

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
	value     string
}

func (e *invalidAttributeTypeError) Error() string {
	return fmt.Sprintf("invalid attribute type: element = %q, attribute = %q, value = %q",
		e.element, e.attribute, e.value)
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

type unimplementedElementError struct {
	elementPath string
}

func (e *unimplementedElementError) Error() string {
	return fmt.Sprintf("unimplemented element: elementPath = %q", e.elementPath)
}

type undefinedConstantError struct {
	element   string
	attribute string
	constant  string
}

func (e *undefinedConstantError) Error() string {
	return fmt.Sprintf("undefined constant: element = %q, attribute = %q, constant = %q",
		e.element, e.attribute, e.constant)
}
