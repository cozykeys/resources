package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshalElementAttributes(e *etree.Element, element *models.KeyboardElementBase) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{
		AttributeName:     {required: false},
		AttributeXOffset:  {required: false},
		AttributeYOffset:  {required: false},
		AttributeRotation: {required: false},
		AttributeHeight:   {required: false},
		AttributeWidth:    {required: false},
		AttributeMargin:   {required: false},
		AttributeDebug:    {required: false},
		AttributeVisible:  {required: false},
	}

	constants := element.GetConstants()

	for _, attr := range e.Attr {
		var err error
		switch attr.Key {
		case AttributeName:
			element.Name, err = unmarshalAttributeString(&attr, constants)
		case AttributeXOffset:
			element.XOffset, err = unmarshalAttributeFloat64(&attr, constants)
		case AttributeYOffset:
			element.YOffset, err = unmarshalAttributeFloat64(&attr, constants)
		case AttributeRotation:
			element.Rotation, err = unmarshalAttributeFloat64(&attr, constants)
		case AttributeHeight:
			element.Height, err = unmarshalAttributeFloat64(&attr, constants)
		case AttributeWidth:
			element.Width, err = unmarshalAttributeFloat64(&attr, constants)
		case AttributeMargin:
			element.Margin, err = unmarshalAttributeFloat64(&attr, constants)
		case AttributeDebug:
			element.Debug, err = unmarshalAttributeBool(&attr, constants)
		case AttributeVisible:
			element.Visible, err = unmarshalAttributeBool(&attr, constants)
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

func isKeyboardElementAttribute(attribute string) bool {
	kbElementAttrs := map[string]interface{}{
		AttributeName:     nil,
		AttributeXOffset:  nil,
		AttributeYOffset:  nil,
		AttributeRotation: nil,
		AttributeHeight:   nil,
		AttributeWidth:    nil,
		AttributeMargin:   nil,
		AttributeDebug:    nil,
		AttributeVisible:  nil,
	}

	_, ok := kbElementAttrs[attribute]
	return ok
}
