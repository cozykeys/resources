package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalSpacer(e *etree.Element, parent models.KeyboardElement) (*models.Spacer, error) {
	unmarshaller := &spacerUnmarshaller{
		element: e,
		parent:  parent,
	}
	return unmarshaller.unmarshal()
}

type spacerUnmarshaller struct {
	element *etree.Element
	spacer  *models.Spacer
	parent  models.KeyboardElement
}

func (u *spacerUnmarshaller) unmarshal() (*models.Spacer, error) {
	if u.element == nil {
		return nil, &nilElementError{}
	}

	if u.element.Tag != ElementSpacer {
		return nil, &invalidTagError{
			expected: ElementSpacer,
			actual:   u.element.Tag,
		}
	}

	u.spacer = &models.Spacer{
		KeyboardElementBase: models.KeyboardElementBase{
			Parent:  u.parent,
			Visible: true,
		},
	}

	if err := findAndUnmarshalConstants(u.element, &u.spacer.KeyboardElementBase); err != nil {
		return nil, err
	}

	err := u.unmarshalAttributes()
	if err != nil {
		return nil, err
	}

	return u.spacer, nil
}

func (u *spacerUnmarshaller) unmarshalAttributes() error {
	return unmarshalElementAttributes(u.element, &u.spacer.KeyboardElementBase)
}
