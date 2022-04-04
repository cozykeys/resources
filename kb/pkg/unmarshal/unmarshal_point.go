package unmarshal

import (
	"kb/pkg/models"
	"strings"

	"github.com/beevik/etree"
)

func unmarshalPoint(e *etree.Element) (*models.Vec2, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	validPointTags := []string{
		"EndPoint",
		"ControlPoint",
	}

	if !stringSliceContains(e.Tag, validPointTags) {
		return nil, &invalidTagError{
			expected: strings.Join(validPointTags, ","),
			actual:   e.Tag,
		}
	}

	endPoint := &models.Vec2{}

	err := unmarshalPointAttributes(endPoint, e.Attr, e.Tag)
	if err != nil {
		return nil, err
	}

	return endPoint, nil
}

func unmarshalPointAttributes(endPoint *models.Vec2, attributes []etree.Attr, tag string) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeX: {required: true},
		AttributeY: {required: true},
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
				element:   tag,
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
				element:   tag,
				attribute: k,
			}
		}
	}

	return nil
}
