package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalStack(e *etree.Element, parent models.KeyboardElement) (*models.Stack, error) {
	unmarshaller := &stackUnmarshaller{
		element: e,
		parent:  parent,
	}

	return unmarshaller.unmarshal()
}

type stackUnmarshaller struct {
	element *etree.Element
	stack   *models.Stack
	parent  models.KeyboardElement
}

func (u *stackUnmarshaller) unmarshal() (*models.Stack, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementStack {
		return nil, &invalidTagError{
			expected: ElementStack,
			actual:   u.element.Tag,
		}
	}

	u.stack = &models.Stack{
		Group: models.Group{
			KeyboardElementBase: models.KeyboardElementBase{
				Parent:  u.parent,
				Visible: true,
			},
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.stack.KeyboardElementBase); err != nil {
		return nil, err
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	err = u.unmarshalChildElements()
	if err != nil {
		return nil, err
	}

	return u.stack, nil
}

func (u *stackUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeOrientation: {required: false},
	}

	constants := u.stack.GetConstants()

	err := unmarshalElementAttributes(u.element, &u.stack.KeyboardElementBase)
	if err != nil {
		return err
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeOrientation:
			u.stack.Orientation, err = unmarshalAttributeStackOrientation(&attr, constants)
		default:
			if !isKeyboardElementAttribute(attr.Key) {
				err = &unexpectedAttributeError{
					element:   ElementStack,
					attribute: attr.Key,
				}
			}
		}

		if err != nil {
			return err
		}

		if a, ok := supportedAttributes[attr.Key]; ok {
			a.found = true
		}
	}

	for k, v := range supportedAttributes {
		if v.required && v.found == false {
			return &missingRequiredAttributeError{
				element:   ElementStack,
				attribute: k,
			}
		}
	}

	return nil
}

func (u *stackUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementChildren:
			u.stack.Children, err = u.unmarshalChildren(element)
		default:
			err = &invalidChildElementError{
				element: ElementStack,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (u *stackUnmarshaller) unmarshalChildren(e *etree.Element) ([]models.GroupChild, error) {
	children := []models.GroupChild{}

	for _, child := range e.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		var child models.GroupChild
		switch element.Tag {
		case ElementGroup:
			child, err = unmarshalGroup(element, u.stack)
		case ElementPath:
			child, err = unmarshalPath(element, u.stack)
		case ElementKey:
			child, err = unmarshalKey(element, u.stack)
		case ElementStack:
			child, err = unmarshalStack(element, u.stack)
		case ElementSpacer:
			child, err = unmarshalSpacer(element, u.stack)
		case ElementCircle:
			child, err = unmarshalCircle(element, u.stack)
		case ElementText:
			child, err = unmarshalText(element, u.stack)
		default:
			return nil, &unimplementedElementError{
				elementPath: getElementPath(element),
			}
		}

		if err != nil {
			return nil, err
		}

		children = append(children, child)
	}

	return children, nil
}
