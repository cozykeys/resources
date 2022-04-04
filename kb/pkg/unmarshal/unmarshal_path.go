package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

/*
   "Path": {
       "Attributes": [
           "Fill",
           "Stroke",
           "FillOpacity",
           "Visible"
       ],
       "Children": [
           "Components"
       ]
   },
*/

func unmarshalPath(e *etree.Element) (*models.Path, error) {
	unmarshaller := newPathUnmarshaller(e)
	return unmarshaller.unmarshal()
}

// TODO: This is the pattern I'd like to move towards for all element
// unmarshalling
type pathUnmarshaller struct {
	element *etree.Element
	path    *models.Path
}

func newPathUnmarshaller(e *etree.Element) *pathUnmarshaller {
	return &pathUnmarshaller{element: e}
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

	u.path = &models.Path{}

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
		AttributeStroke:      {required: false},
		AttributeFill:        {required: false},
		AttributeFillOpacity: {required: false},
		AttributeVisible:     {required: false},
	}

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeFill:
			u.path.Fill, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeFillOpacity:
			u.path.FillOpacity, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeStroke:
			u.path.Stroke, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeStrokeWidth:
			u.path.StrokeWidth, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeVisible:
			u.path.Visible, err = unmarshalAttributeBool(attr.Key, attr.Value)
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

func (u *pathUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementComponents:
			u.path.Components, err = unmarshalPathComponents(element)
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
			// TODO: Implement the other path components
			/*
				case ElementAbsoluteCubicCurveTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementAbsoluteHorizontalLineTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementAbsoluteQuadraticCurveTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementAbsoluteVerticalLineTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementRelativeCubicCurveTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementRelativeHorizontalLineTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementRelativeLineTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementRelativeMoveTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementRelativeQuadraticCurveTo:
					pathComponent, err = unmarshalFoo(element)
				case ElementRelativeVerticalLineTo:
					pathComponent, err = unmarshalFoo(element)
			*/
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
