package unmarshal

import (
	"kb/pkg/models"
	"strings"

	"github.com/beevik/etree"
)

func unmarshalPoint(e *etree.Element, parent models.KeyboardElement) (*models.Point, error) {
	unmarshaller := &pointUnmarshaller{
		element: e,
		parent:  parent,
	}

	return unmarshaller.unmarshal()
}

type pointUnmarshaller struct {
	element *etree.Element
	point   *models.Point
	parent  models.KeyboardElement
}

func (u *pointUnmarshaller) unmarshal() (*models.Point, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	validPointTags := []string{
		"EndPoint",
		"ControlPoint",
	}

	if !stringSliceContains(u.element.Tag, validPointTags) {
		return nil, &invalidTagError{
			expected: strings.Join(validPointTags, ","),
			actual:   getElementPath(u.element),
		}
	}

	u.point = &models.Point{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent: u.parent,
		},
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	return u.point, nil
}

func (u *pointUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeX: {required: true},
		AttributeY: {required: true},
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeX:
			u.point.X, err = unmarshalAttributeFloat64(&attr, u.point.GetConstants())
		case AttributeY:
			u.point.Y, err = unmarshalAttributeFloat64(&attr, u.point.GetConstants())
		default:
			err = &unexpectedAttributeError{
				element:   getElementPath(u.element),
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
				element:   getElementPath(u.element),
				attribute: k,
			}
		}
	}

	return nil
}
