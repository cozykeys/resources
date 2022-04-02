package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

// TODO: Change this file and symbol names from EndPoint to Vec2 or Point or
// something
func unmarshalEndPoint(e *etree.Element) (*models.Vec2, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementEndPoint {
		return nil, &invalidTagError{
			expected: ElementEndPoint,
			actual:   e.Tag,
		}
	}

	endPoint := &models.Vec2{}

	err := unmarshalEndPointAttributes(endPoint, e.Attr)
	if err != nil {
		return nil, err
	}

	return endPoint, nil
}

func unmarshalEndPointAttributes(endPoint *models.Vec2, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeX: {required: false},
		AttributeY: {required: false},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeX:
			endPoint.X, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeY:
			endPoint.Y, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		default:
			err = &unexpectedAttributeError{
				element:   ElementEndPoint,
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

	for k, v := range supportedAttributes {
		if v.required && v.found == false {
			return &missingRequiredAttributeError{
				element:   ElementEndPoint,
				attribute: k,
			}
		}
	}

	return nil
}
