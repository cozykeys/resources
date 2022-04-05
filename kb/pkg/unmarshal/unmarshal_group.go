package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalGroups(e *etree.Element, parent models.KeyboardElement) ([]models.Group, error) {
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

		group, err := unmarshalGroup(element, parent)
		if err != nil {
			return nil, err
		}

		groups = append(groups, *group)
	}

	return groups, nil
}

func unmarshalGroup(e *etree.Element, parent models.KeyboardElement) (*models.Group, error) {
	unmarshaller := &groupUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type groupUnmarshaller struct {
	element *etree.Element
	group   *models.Group
	parent  models.KeyboardElement
}

func (u *groupUnmarshaller) unmarshal() (*models.Group, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementGroup {
		return nil, &invalidTagError{
			expected: ElementGroup,
			actual:   u.element.Tag,
		}
	}

	u.group = &models.Group{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent: u.parent,
		},
	}

	if err := u.unmarshalConstants(); err != nil {
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

	return u.group, nil
}

func (u *groupUnmarshaller) unmarshalAttributes() error {
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

	for _, attr := range u.element.Attr {
		var err error
		switch attr.Key {
		case AttributeName:
			u.group.Name, err = unmarshalAttributeString(&attr, u.group.GetConstants())
		case AttributeRotation:
			u.group.Rotation, err = unmarshalAttributeFloat64(&attr, u.group.GetConstants())
		case AttributeXOffset:
			u.group.XOffset, err = unmarshalAttributeFloat64(&attr, u.group.GetConstants())
		case AttributeYOffset:
			u.group.YOffset, err = unmarshalAttributeFloat64(&attr, u.group.GetConstants())
		case AttributeVisible:
			u.group.Visible, err = unmarshalAttributeBool(&attr, u.group.GetConstants())
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

func (u *groupUnmarshaller) unmarshalChildElements() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		var err error
		switch element.Tag {
		case ElementChildren:
			u.group.Children, err = u.unmarshalChildren(element)
		case ElementConstants:
			u.group.Constants, err = unmarshalConstants(element, u.group)
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

func (u *groupUnmarshaller) unmarshalChildren(e *etree.Element) ([]models.GroupChild, error) {
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
			child, err = unmarshalGroup(element, u.group)
		case ElementPath:
			child, err = unmarshalPath(element, u.group)
		case ElementKey:
			child, err = unmarshalKey(element, u.group)
		case ElementStack:
			child, err = unmarshalStack(element, u.group)
		case ElementSpacer:
			child, err = unmarshalSpacer(element, u.group)
		case ElementCircle:
			child, err = unmarshalCircle(element, u.group)
		case ElementText:
			child, err = unmarshalText(element, u.group)
		default:
			return nil, &unimplementedElementError{
				elementPath: getElementPath(element),
			}
		}

		if err != nil {
			return nil, err
		}

		children = append(children, child)
	}

	return children, nil
}

func (u *groupUnmarshaller) unmarshalConstants() error {
	for _, child := range u.element.Child {
		element, ok := child.(*etree.Element)
		if !ok {
			continue
		}

		if element.Tag == ElementConstants {
			var err error
			u.group.Constants, err = unmarshalConstants(element, u.group)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}
