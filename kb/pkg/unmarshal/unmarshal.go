package unmarshal

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"kb/pkg/models"

	"github.com/beevik/etree"
)

const (
	ElementKeyboard  = "Keyboard"
	ElementLayers    = "Layers"
	ElementConstants = "Constants"
	ElementConstant  = "Constant"

	AttributeName    = "Name"
	AttributeVersion = "Version"
	AttributeValue   = "Value"
)

func Unmarshal(bytes []byte) (*models.Keyboard, error) {
	doc := etree.NewDocument()

	err := doc.ReadFromBytes(bytes)
	if err != nil {
		return nil, err
	}

	keyboard, err := unmarshalKeyboard(doc.Root())
	if err != nil {
		return nil, err
	}

	return keyboard, nil
}

func unmarshalAttributeString(key, raw string) (string, error) {
	// TODO: Process constants
	return raw, nil
}

func unmarshalAttributeFloat64(key, raw string) (float64, error) {
	// TODO: Process constants
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, &invalidAttributeTypeError{
			element:   ElementKeyboard,
			attribute: key,
		}
	}
	return val, nil
}

type XmlMeta struct {
	Attributes []string
	Children   []string
}

// TODO: Temporary code, delete this
func WalkTree(bytes []byte) {
	doc := etree.NewDocument()

	err := doc.ReadFromBytes(bytes)
	if err != nil {
		panic(err)
	}

	xmlMeta := map[string]*XmlMeta{}

	walkTree(doc.Root(), xmlMeta)

	jsonBytes, err := json.Marshal(xmlMeta)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))
}

// TODO: Temporary code, delete this
func walkTree(e *etree.Element, xmlMetaMap map[string]*XmlMeta) {
	if e == nil {
		panic("element is nil")
	}

	xmlMeta, ok := xmlMetaMap[e.Tag]
	if !ok {
		xmlMeta = &XmlMeta{
			Attributes: []string{},
			Children:   []string{},
		}
		xmlMetaMap[e.Tag] = xmlMeta
	}

	for _, attr := range e.Attr {
		if !stringSliceContains(attr.Key, xmlMeta.Attributes) {
			xmlMeta.Attributes = append(xmlMeta.Attributes, attr.Key)
		}
	}

	for _, child := range e.Child {
		var childElement *etree.Element

		switch v := child.(type) {
		case *etree.Element:
			childElement = v
		case *etree.CharData:
			log.Printf("Skipping child of type CharData, Data = %q", v.Data)
			continue
		case *etree.Comment:
			fmt.Println("Skipping child of type Comment")
			continue
		case *etree.Directive:
			fmt.Println("Skipping child of type Directive")
			continue
		case *etree.ProcInst:
			fmt.Println("Skipping child of type ProcInst")
			continue
		default:
			panic("unknown type")
		}

		if !stringSliceContains(childElement.Tag, xmlMeta.Children) {
			xmlMeta.Children = append(xmlMeta.Children, childElement.Tag)
		}

		walkTree(childElement, xmlMetaMap)
	}
}

func stringSliceContains(needle string, haystack []string) bool {
	for _, s := range haystack {
		if needle == s {
			return true
		}
	}

	return false
}
