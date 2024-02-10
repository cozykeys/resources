package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalPath(e *etree.Element, parent models.KeyboardElement) (*models.Path, error) {
	unmarshaller := &pathUnmarshaller{
		element: e,
		parent:  parent,
	}

	return unmarshaller.unmarshal()
}

type pathUnmarshaller struct {
	element *etree.Element
	path    *models.Path
	parent  models.KeyboardElement
}

func (u *pathUnmarshaller) unmarshal() (*models.Path, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementPath {
		return nil, &invalidTagError{
			expected: ElementPath,
			actual:   u.element.Tag,
		}
	}

	u.path = &models.Path{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.path.KeyboardElementBase); err != nil {
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

	return u.path, nil
}

func (u *pathUnmarshaller) unmarshalAttributes() error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeStrokeWidth: {required: false},
		AttributeStroke:      {required: false},
		AttributeFill:        {required: false},
		AttributeFillOpacity: {required: false},
	}

	constants := u.path.GetConstants()

	err := unmarshalElementAttributes(u.element, &u.path.KeyboardElementBase)
	if err != nil {
		return err
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeFill:
			u.path.Fill, err = unmarshalAttributeString(&attr, constants)
		case AttributeFillOpacity:
			u.path.FillOpacity, err = unmarshalAttributeString(&attr, constants)
		case AttributeStroke:
			u.path.Stroke, err = unmarshalAttributeString(&attr, constants)
		case AttributeStrokeWidth:
			u.path.StrokeWidth, err = unmarshalAttributeString(&attr, constants)
		default:
			if !isKeyboardElementAttribute(attr.Key) {
				err = &unexpectedAttributeError{
					element:   ElementPath,
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
				element:   ElementPath,
				attribute: k,
			}
		}
	}

	return nil
}

func (u *pathUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementComponents:
			u.path.Components, err = u.unmarshalPathComponents(element)
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

func (u *pathUnmarshaller) unmarshalPathComponents(e *etree.Element) ([]models.PathComponent, error) {
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
			pathComponent, err = unmarshalAbsoluteMoveTo(element, u.path)
		case ElementAbsoluteLineTo:
			pathComponent, err = unmarshalAbsoluteLineTo(element, u.path)
		case ElementAbsoluteQuadraticCurveTo:
			pathComponent, err = unmarshalAbsoluteQuadraticCurveTo(element, u.path)
			// TODO: Implement the other path components
			/*
				case ElementAbsoluteCubicCurveTo:
					pathComponent, err = unmarshalAbsoluteCubicCurveTo(element, u.path)
				case ElementAbsoluteHorizontalLineTo:
					pathComponent, err = unmarshalAbsoluteHorizontalLineTo(element, u.path)
				case ElementAbsoluteVerticalLineTo:
					pathComponent, err = unmarshalAbsoluteVerticalLineTo(element, u.path)
				case ElementRelativeCubicCurveTo:
					pathComponent, err = unmarshalRelativeCubicCurveTo(element, u.path)
				case ElementRelativeHorizontalLineTo:
					pathComponent, err = unmarshalRelativeHorizontalLineTo(element, u.path)
				case ElementRelativeLineTo:
					pathComponent, err = unmarshalRelativeLineTo(element, u.path)
				case ElementRelativeMoveTo:
					pathComponent, err = unmarshalRelativeMoveTo(element, u.path)
				case ElementRelativeQuadraticCurveTo:
					pathComponent, err = unmarshalRelativeQuadraticCurveTo(element, u.path)
				case ElementRelativeVerticalLineTo:
					pathComponent, err = unmarshalRelativeVerticalLineTo(element, u.path)
			*/
		default:
			return nil, &invalidChildElementError{
				element: e.Tag,
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
