package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalConstants(e *etree.Element) ([]models.Constant, error) {
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

		constant, err := unmarshalConstant(element)
		if err != nil {
			return nil, err
		}

		constants = append(constants, *constant)
	}

	return constants, nil
}

func unmarshalConstant(e *etree.Element) (*models.Constant, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementConstant {
		return nil, &invalidTagError{
			expected: ElementConstant,
			actual:   e.Tag,
		}
	}

	constant := &models.Constant{}

	err := unmarshalConstantAttributes(constant, e.Attr)
	if err != nil {
		return nil, err
	}

	return constant, nil
}

func unmarshalConstantAttributes(constant *models.Constant, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeName:  {required: true},
		AttributeValue: {required: true},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeName:
			constant.Name, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeValue:
			constant.Value, err = unmarshalAttributeString(attr.Key, attr.Value)
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
