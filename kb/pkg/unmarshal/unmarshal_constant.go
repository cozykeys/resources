package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalConstants(e *etree.Element, parent models.KeyboardElement) ([]models.Constant, error) {
	constants := []models.Constant{}

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

		constants = append(constants, *constant)
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
			Parent: u.parent,
		},
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
		AttributeName:  {required: true},
		AttributeValue: {required: true},
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeName:
			u.constant.Name, err = unmarshalAttributeString(&attr, u.constant.GetConstants())
		case AttributeValue:
			u.constant.Value, err = unmarshalAttributeString(&attr, u.constant.GetConstants())
		default:
			err = &unexpectedAttributeError{
				element:   ElementConstant,
				attribute: attr.Key,
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
