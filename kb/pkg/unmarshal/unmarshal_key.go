package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalKey(e *etree.Element) (*models.Key, error) {
	unmarshaller := newKeyUnmarshaller(e)
	return unmarshaller.unmarshal()
}

// TODO: This is the pattern I'd like to move towards for all element
// unmarshalling
type keyUnmarshaller struct {
	element *etree.Element
	key     *models.Key
}

func newKeyUnmarshaller(e *etree.Element) *keyUnmarshaller {
	return &keyUnmarshaller{element: e}
}

func (u *keyUnmarshaller) unmarshal() (*models.Key, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementKey {
		return nil, &invalidTagError{
			expected: ElementKey,
			actual:   u.element.Tag,
		}
	}

	u.key = &models.Key{}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	err = u.unmarshalChildElements()
	if err != nil {
		return nil, err
	}

	return u.key, nil
}

func (u *keyUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeName:    {required: false},
		AttributeRow:     {required: false},
		AttributeColumn:  {required: false},
		AttributeXOffset: {required: false},
		AttributeYOffset: {required: false},
		AttributeWidth:   {required: false},
		AttributeHeight:  {required: false},
		AttributeMargin:  {required: false},
		AttributeFill:    {required: false},
		AttributeStroke:  {required: false},
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeName:
			u.key.Name, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeRow:
			u.key.Row, err = unmarshalAttributeInt(attr.Key, attr.Value)
		case AttributeColumn:
			u.key.Column, err = unmarshalAttributeInt(attr.Key, attr.Value)
		case AttributeXOffset:
			u.key.XOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeYOffset:
			u.key.YOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeWidth:
			u.key.Width, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeHeight:
			u.key.Height, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeMargin:
			u.key.Margin, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeFill:
			u.key.Fill, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeStroke:
			u.key.Stroke, err = unmarshalAttributeString(attr.Key, attr.Value)
		default:
			err = &unexpectedAttributeError{
				element:   ElementKey,
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
				element:   ElementKey,
				attribute: k,
			}
		}
	}

	return nil
}

func (u *keyUnmarshaller) unmarshalChildElements() error {
	legends := []models.Legend{}
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		// TODO: Put <Legend /> elements in a <Legends /> element so we can
		// change this to:
		// case ElementLegends:
		//     u.key.Legends, err = unmarshalLegends(element)
		case ElementLegend:
			var legend *models.Legend
			legend, err = unmarshalLegend(element)
			legends = append(legends, *legend)
		default:
			err = &invalidChildElementError{
				element: ElementKey,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	u.key.Legends = legends
	return nil
}
