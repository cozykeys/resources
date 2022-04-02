package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalGroups(e *etree.Element) ([]models.Group, error) {
	groups := []models.Group{}

	for _, child := range e.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		if element.Tag != ElementGroup {
			return nil, &invalidChildElementError{
				element: ElementGroups,
				child:   element.Tag,
			}
		}

		group, err := unmarshalGroup(element)
		if err != nil {
			return nil, err
		}

		groups = append(groups, *group)
	}

	return groups, nil
}

func unmarshalGroup(e *etree.Element) (*models.Group, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != ElementGroup {
		return nil, &invalidTagError{
			expected: ElementGroup,
			actual:   e.Tag,
		}
	}

	group := &models.Group{}

	err := unmarshalGroupAttributes(group, e.Attr)
	if err != nil {
		return nil, err
	}

	err = unmarshalGroupChildren(group, e.Child)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func unmarshalGroupAttributes(group *models.Group, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeName:     {required: true},
		AttributeRotation: {required: false},
		AttributeXOffset:  {required: false},
		AttributeYOffset:  {required: false},
		AttributeVisible:  {required: false},
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key {
		case AttributeName:
			group.Name, err = unmarshalAttributeString(attr.Key, attr.Value)
		case AttributeRotation:
			group.Rotation, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeXOffset:
			group.XOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeYOffset:
			group.YOffset, err = unmarshalAttributeFloat64(attr.Key, attr.Value)
		case AttributeVisible:
			group.Visible, err = unmarshalAttributeBool(attr.Key, attr.Value)
		default:
			err = &unexpectedAttributeError{
				element:   ElementGroup,
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
				element:   ElementGroup,
				attribute: attrKey,
			}
		}
	}

	return nil
}

func unmarshalGroupChildren(group *models.Group, children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementChildren:
			group.Children, err = unmarshalChildren(element)
		default:
			err = &invalidChildElementError{
				element: ElementGroup,
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// TODO: Move this into unmarshal_children.go?
func unmarshalChildren(e *etree.Element) ([]models.GroupChild, error) {
	children := []models.GroupChild{}

	for _, child := range e.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		var child models.GroupChild
		switch element.Tag {
		case ElementGroup:
			child, err = unmarshalGroup(element)
		case ElementPath:
			// TODO
			//child, err = unmarshalPath(element)
		case ElementKey:
			// TODO
			//child, err = unmarshalKey(element)
		default:
			return nil, &invalidChildElementError{
				element: ElementGroups,
				child:   element.Tag,
			}
		}

		if err != nil {
			return nil, err
		}

		children = append(children, child)
	}

	return children, nil
}
