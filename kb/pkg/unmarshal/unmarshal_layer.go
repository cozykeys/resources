package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

/*
   "Layer": {
       "Attributes": [
           "Name",
           "ZIndex",
           "XOffset",
           "YOffset",
           "Width",
           "Height"
       ],
       "Children": [
           "Groups"
       ]
   },
*/

func unmarshalLayers(e *etree.Element) ([]models.Layer, error) {
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

		layer, err := unmarshalLayer(element)
		if err != nil {
			return nil, err
		}

		layers = append(layers, *layer)
	}

	return layers, nil
}

func unmarshalLayer(e *etree.Element) (*models.Layer, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementLayer {
		return nil, &invalidTagError{
			expected: ElementLayer,
			actual:   e.Tag,
		}
	}

	layer := &models.Layer{}

	err := unmarshalLayerAttributes(layer, e.Attr)
	if err != nil {
		return nil, err
	}

	err = unmarshalLayerChildElements(layer, e.Child)
	if err != nil {
		return nil, err
	}

	return layer, nil
}

func unmarshalLayerAttributes(layer *models.Layer, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeName:    {required: true},
		AttributeZIndex:  {required: false},
		AttributeXOffset: {required: false},
		AttributeYOffset: {required: false},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeName:
			layer.Name, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeZIndex:
			layer.ZIndex, err = unmarshalAttributeInt(attr.Key, attr.Value)
		case AttributeXOffset:
			layer.XOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeYOffset:
			layer.YOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		default:
			err = &unexpectedAttributeError{
				element:   ElementLayer,
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
				element:   ElementLayer,
				attribute: attrKey,
			}
		}
	}

	return nil
}

func unmarshalLayerChildElements(layer *models.Layer, children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementGroups:
			layer.Groups, err = unmarshalGroups(element)
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
