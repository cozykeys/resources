package kb

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/beevik/etree"
)

func Unmarshal(bytes []byte) (*Keyboard, error) {
	doc := etree.NewDocument()

	err := doc.ReadFromBytes(bytes)
	if err != nil {
		return nil, err
	}

	keyboard := &Keyboard{}
	err = keyboard.unmarshal(doc.Root())
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
	child *XmlMeta
}

// TODO: Temporary code, delete this
func WalkTree(bytes []byte) {
	doc := etree.NewDocument()

	err := doc.ReadFromBytes(bytes)
	if err != nil {
		panic(err)
	}

	attrMap := make(map[string][]string)
	childMap := make(map[string][]string)

	walkTree(doc.Root(), attrMap, childMap)

	attrMapJSON, err := json.Marshal(attrMap)
	if err != nil {
		panic(err)
	}

	childMapJSON, err := json.Marshal(childMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(attrMapJSON))
	fmt.Println(string(childMapJSON))
}

// TODO: Temporary code, delete this
func walkTree(e *etree.Element, attrMap, childMap map[string][]string) {
	if e == nil {
		panic("element is nil")
	}

	for _, attr := range e.Attr {
		attrMapEntry, ok := attrMap[e.Tag]
		if !ok {
			attrMap[e.Tag] = []string{}
		}

		attrMapEntry, _ = attrMap[e.Tag]

		if stringSliceContains(attr.Key, attrMapEntry) {
			continue
		}

		attrMapEntry = append(attrMapEntry, attr.Key)
		attrMap[e.Tag] = attrMapEntry
	}

	for _, child := range e.Child {
		childElement, ok := child.(*etree.Element)

		switch v := child.(type) {
		case *etree.Element:
			// Do nothing
		case *etree.CharData:
			log.Printf("Skipping child of type CharData, Data = %q", v.Data)
		case *etree.Comment:
			fmt.Println("Skipping child of type Comment")
		case *etree.Directive:
			fmt.Println("Skipping child of type Directive")
		case *etree.ProcInst:
			fmt.Println("Skipping child of type ProcInst")
		default:
			panic("unknown type")
		}
		if !ok {
			continue
		}

		childMapEntry, ok := childMap[e.Tag]
		if !ok {
			childMap[e.Tag] = []string{}
		}

		childMapEntry, _ = childMap[e.Tag]

		if stringSliceContains(childElement.Tag, childMapEntry) {
			continue
		}

		childMapEntry = append(childMapEntry, childElement.Tag)
		childMap[e.Tag] = childMapEntry

		walkTree(childElement, attrMap, childMap)
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
