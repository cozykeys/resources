package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalLayers(e *etree.Element, parent models.KeyboardElement) ([]models.Layer, error) {
	layers := []models.Layer{}

	for _, child := range e.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		if element.Tag != ElementLayer {
			return nil, &invalidChildElementError{
				element: ElementLayers,
				child:   element.Tag,
			}
		}

		layer, err := unmarshalLayer(element, parent)
		if err != nil {
			return nil, err
		}

		layers = append(layers, *layer)
	}

	return layers, nil
}

func unmarshalLayer(e *etree.Element, parent models.KeyboardElement) (*models.Layer, error) {
	unmarshaller := &layerUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type layerUnmarshaller struct {
	element *etree.Element
	layer   *models.Layer
	parent  models.KeyboardElement
}

func (u *layerUnmarshaller) unmarshal() (*models.Layer, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementLayer {
		return nil, &invalidTagError{
			expected: ElementLayer,
			actual:   u.element.Tag,
		}
	}

	u.layer = &models.Layer{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.layer.KeyboardElementBase); err != nil {
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

	return u.layer, nil
}

func (u *layerUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeZIndex: {required: false},
	}

	constants := u.layer.GetConstants()

	err := unmarshalElementAttributes(u.element, &u.layer.KeyboardElementBase)
	if err != nil {
		return err
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeZIndex:
			u.layer.ZIndex, err = unmarshalAttributeInt(&attr, constants)
		default:
			if !isKeyboardElementAttribute(attr.Key) {
				err = &unexpectedAttributeError{
					element:   ElementLayer,
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
				element:   ElementLayer,
				attribute: attrKey,
			}
		}
	}

	return nil
}

func (u *layerUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementGroups:
			u.layer.Groups, err = unmarshalGroups(element, u.layer)
		default:
			err = &invalidChildElementError{
				element: ElementLayer,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
