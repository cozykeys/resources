package unmarshal

import (
	"kb/pkg/models"
	"log"

	"github.com/beevik/etree"
)

// TODO Move this to unmarshal_constant.go
func unmarshalConstants(e *etree.Element) ([]models.Constant, error) {
	// TODO: Implement this
	log.Print("unmarshalConstants")
	return nil, nil
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

	kb := &models.Constant{}

	err := unmarshalConstantAttributes(kb, e.Attr)
	if err != nil {
		return nil, err
	}

	return kb, nil
}

func unmarshalConstantAttributes(kb *models.Constant, attributes []etree.Attr) error {
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
			kb.Name, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeValue:
			kb.Value, err = unmarshalAttributeString(attr.Key, attr.Value)
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
