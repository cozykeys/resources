package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalKey(e *etree.Element, parent models.KeyboardElement) (*models.Key, error) {
	unmarshaller := &keyUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type keyUnmarshaller struct {
	element *etree.Element
	key     *models.Key
	parent  models.KeyboardElement
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

	u.key = &models.Key{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.key.KeyboardElementBase); err != nil {
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

	return u.key, nil
}

func (u *keyUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeRow:    {required: false},
		AttributeColumn: {required: false},
		AttributeFill:   {required: false},
		AttributeStroke: {required: false},
	}

	constants := u.key.GetConstants()

	err := unmarshalElementAttributes(u.element, &u.key.KeyboardElementBase)
	if err != nil {
		return err
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeRow:
			u.key.Row, err = unmarshalAttributeInt(&attr, constants)
		case AttributeColumn:
			u.key.Column, err = unmarshalAttributeInt(&attr, constants)
		case AttributeFill:
			u.key.Fill, err = unmarshalAttributeString(&attr, constants)
		case AttributeStroke:
			u.key.Stroke, err = unmarshalAttributeString(&attr, constants)
		default:
			if !isKeyboardElementAttribute(attr.Key) {
				err = &unexpectedAttributeError{
					element:   ElementKey,
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
			legend, err = unmarshalLegend(element, u.key)
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
