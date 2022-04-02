package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalPath(e *etree.Element) (*models.Path, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementPath {
		return nil, &invalidTagError{
			expected: ElementPath,
			actual:   e.Tag,
		}
	}

	path := &models.Path{}

	err := unmarshalPathAttributes(path, e.Attr)
	if err != nil {
		return nil, err
	}

	err = unmarshalPathChildren(path, e.Child)
	if err != nil {
		return nil, err
	}

	return path, nil
}

func unmarshalPathAttributes(path *models.Path, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeStroke:      {required: false},
		AttributeFill:        {required: false},
		AttributeFillOpacity: {required: false},
		AttributeVisible:     {required: false},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeFill:
			path.Fill, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeFillOpacity:
			path.FillOpacity, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeStroke:
			path.Stroke, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeStrokeWidth:
			path.StrokeWidth, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeVisible:
			path.Visible, err = unmarshalAttributeBool(attr.Key, attr.Value)
		default:
			err = &unexpectedAttributeError{
				element:   ElementPath,
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
				element:   ElementPath,
				attribute: k,
			}
		}
	}

	return nil
}

func unmarshalPathChildren(path *models.Path, children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementComponents:
			path.Components, err = unmarshalPathComponents(element)
		default:
			err = &invalidChildElementError{
				element: ElementPath,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func unmarshalPathComponents(e *etree.Element) ([]models.PathComponent, error) {
	pathComponents := []models.PathComponent{}

	for _, child := range e.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		var pathComponent models.PathComponent
		switch element.Tag {
		case ElementAbsoluteMoveTo:
			pathComponent, err = unmarshalAbsoluteMoveTo(element)
		case ElementAbsoluteLineTo:
			pathComponent, err = unmarshalAbsoluteLineTo(element)
		default:
			return nil, &invalidChildElementError{
				element: ElementGroups,
				child:   element.Tag,
			}
		}

		if err != nil {
			return nil, err
		}

		pathComponents = append(pathComponents, pathComponent)
	}

	return pathComponents, nil
}
