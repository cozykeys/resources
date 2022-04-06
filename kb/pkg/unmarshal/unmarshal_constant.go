package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func findAndUnmarshalConstants(e *etree.Element, kbElement *models.KeyboardElementBase) error {
	for _, child := range e.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		if element.Tag == ElementConstants {
			var err error
			kbElement.Constants, err = unmarshalConstants(element, kbElement)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}

func unmarshalConstants(e *etree.Element, parent models.KeyboardElement) (map[string]string, error) {
	constants := map[string]string{}

	for _, child := range e.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		if element.Tag != ElementConstant {
			return nil, &invalidChildElementError{
				element: ElementConstants,
				child:   element.Tag,
			}
		}

		constant, err := unmarshalConstant(element, parent)
		if err != nil {
			return nil, err
		}

		constants[constant.Name] = constant.Value
	}

	return constants, nil
}

func unmarshalConstant(e *etree.Element, parent models.KeyboardElement) (*models.Constant, error) {
	unmarshaller := &constantUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type constantUnmarshaller struct {
	element  *etree.Element
	constant *models.Constant
	parent   models.KeyboardElement
}

func (u *constantUnmarshaller) unmarshal() (*models.Constant, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementConstant {
		return nil, &invalidTagError{
			expected: ElementConstant,
			actual:   u.element.Tag,
		}
	}

	u.constant = &models.Constant{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.constant.KeyboardElementBase); err != nil {
		return nil, err
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	return u.constant, nil
}

func (u *constantUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeValue: {required: true},
	}

	constants := u.constant.GetConstants()

	err := unmarshalElementAttributes(u.element, &u.constant.KeyboardElementBase)
	if err != nil {
		return err
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeValue:
			u.constant.Value, err = unmarshalAttributeString(&attr, constants)
		default:
			if !isKeyboardElementAttribute(attr.Key) {
				err = &unexpectedAttributeError{
					element:   ElementConstant,
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

	for attrKey, attrVal := range supportedAttributes {
		if attrVal.required && attrVal.found == false {
			return &missingRequiredAttributeError{
				element:   ElementConstant,
				attribute: attrKey,
			}
		}
	}

	return nil
}
