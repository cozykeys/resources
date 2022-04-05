package unmarshal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"kb/pkg/models"

	"github.com/beevik/etree"
)

const (
	ElementAbsoluteLineTo = "AbsoluteLineTo"
	ElementAbsoluteMoveTo = "AbsoluteMoveTo"
	ElementChildren       = "Children"
	ElementCircle         = "Circle"
	ElementComponents     = "Components"
	ElementConstant       = "Constant"
	ElementConstants      = "Constants"
	ElementEndPoint       = "EndPoint"
	ElementGroup          = "Group"
	ElementGroups         = "Groups"
	ElementKey            = "Key"
	ElementKeyboard       = "Keyboard"
	ElementLayer          = "Layer"
	ElementLayers         = "Layers"
	ElementLegend         = "Legend"
	ElementPath           = "Path"
	ElementSpacer         = "Spacer"
	ElementStack          = "Stack"
	ElementText           = "Text"

	AttributeColor               = "Color"
	AttributeColumn              = "Column"
	AttributeContent             = "Content"
	AttributeFill                = "Fill"
	AttributeFillOpacity         = "FillOpacity"
	AttributeFont                = "Font"
	AttributeFontSize            = "FontSize"
	AttributeHeight              = "Height"
	AttributeHorizontalAlignment = "HorizontalAlignment"
	AttributeMargin              = "Margin"
	AttributeName                = "Name"
	AttributeOrientation         = "Orientation"
	AttributeRotation            = "Rotation"
	AttributeRow                 = "Row"
	AttributeSize                = "Size"
	AttributeStroke              = "Stroke"
	AttributeStrokeWidth         = "StrokeWidth"
	AttributeText                = "Text"
	AttributeTextAnchor          = "TextAnchor"
	AttributeValue               = "Value"
	AttributeVersion             = "Version"
	AttributeVerticalAlignment   = "VerticalAlignment"
	AttributeVisible             = "Visible"
	AttributeWidth               = "Width"
	AttributeX                   = "X"
	AttributeXOffset             = "XOffset"
	AttributeY                   = "Y"
	AttributeYOffset             = "YOffset"
	AttributeZIndex              = "ZIndex"
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

func unmarshalAttributeString(attr *etree.Attr, constants []models.Constant) (string, error) {
	value, err := expandConstants(attr, constants)
	if err != nil {
		return "", err
	}
	return value, nil
}

func unmarshalAttributeBool(attr *etree.Attr, constants []models.Constant) (bool, error) {
	value, err := expandConstants(attr, constants)
	if err != nil {
		return false, err
	}

	lower := strings.ToLower(value)
	switch lower {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		// TODO: This error is a bit confusing currently; make a new
		// "invalidAttributeValueError" instead
		return false, &invalidAttributeTypeError{
			element:   attr.Element().Tag,
			attribute: attr.Key,
		}
	}
}

func unmarshalAttributeFloat64(attr *etree.Attr, constants []models.Constant) (float64, error) {
	value, err := expandConstants(attr, constants)
	if err != nil {
		return 0.0, err
	}

	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, &invalidAttributeTypeError{
			element:   attr.Element().Tag,
			attribute: attr.Key,
			value:     attr.Value,
		}
	}
	return val, nil
}

func unmarshalAttributeInt(attr *etree.Attr, constants []models.Constant) (int, error) {
	value, err := expandConstants(attr, constants)
	if err != nil {
		return 0, err
	}

	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, &invalidAttributeTypeError{
			element:   attr.Element().Tag,
			attribute: attr.Key,
		}
	}
	return int(val), nil
}

func unmarshalAttributeLegendHorizontalAlignment(attr *etree.Attr, constants []models.Constant) (models.LegendHorizontalAlignment, error) {
	value, err := expandConstants(attr, constants)
	if err != nil {
		return 0, err
	}

	val, ok := models.LegendHorizontalAlignmentStr[value]
	if !ok {
		return models.LegendHorizontalAlignmentLeft, &invalidAttributeTypeError{
			element:   attr.Element().Tag,
			attribute: attr.Key,
		}
	}

	return val, nil
}

func unmarshalAttributeLegendVerticalAlignment(attr *etree.Attr, constants []models.Constant) (models.LegendVerticalAlignment, error) {
	value, err := expandConstants(attr, constants)
	if err != nil {
		return 0, err
	}

	val, ok := models.LegendVerticalAlignmentStr[value]
	if !ok {
		return models.LegendVerticalAlignmentTop, &invalidAttributeTypeError{
			element:   attr.Element().Tag,
			attribute: attr.Key,
		}
	}

	return val, nil
}

func unmarshalAttributeStackOrientation(attr *etree.Attr, constants []models.Constant) (models.StackOrientation, error) {
	value, err := expandConstants(attr, constants)
	if err != nil {
		return 0, err
	}

	val, ok := models.StackOrientationStr[value]
	if !ok {
		return models.StackOrientationHorizontal, &invalidAttributeTypeError{
			element:   attr.Element().Tag,
			attribute: attr.Key,
		}
	}

	return val, nil
}

func expandConstants(attr *etree.Attr, constants []models.Constant) (string, error) {
	if constants == nil && len(constants) < 1 {
		return attr.Value, nil
	}

	re := regexp.MustCompile(`\$\{([a-zA-Z0-9]+)\}`)
	constantsFound := re.FindAllString(attr.Value, -1)

	expanded := attr.Value
	for _, c := range constantsFound {
		constantName := strings.Trim(c, "${}")
		constant, ok := getConstant(constantName, constants)
		if !ok {
			return "", &undefinedConstantError{
				element:   getElementPath(attr.Element()),
				attribute: attr.Key,
				constant:  constantName,
			}
		}
		expanded = strings.ReplaceAll(expanded, c, constant.Value)
	}
	return expanded, nil
}

func getConstant(name string, constants []models.Constant) (*models.Constant, bool) {
	for _, c := range constants {
		if name == c.Name {
			return &c, true
		}
	}
	return nil, false
}

func expandConstants2() {
	constants := map[string]string{
		"VersionMajor": "1",
		"VersionMinor": "2",
	}

	s := "v${VersionMajor}.${VersionMinor}"
	re := regexp.MustCompile(`\$\{([a-zA-Z]+)\}`)
	constantsFound := re.FindAllString(s, -1)
	for _, c := range constantsFound {
		constantName := strings.Trim(c, "${}")
		constant, ok := constants[constantName]
		if !ok {
			panic("TODO")
		}
		s = strings.ReplaceAll(s, c, constant)
	}
}

// TODO: Temporary code, delete this
type XmlMeta struct {
	Attributes []string
	Children   []string
}

// TODO: Temporary code, delete this
func WalkFiles(files []string) {
	xmlMeta := map[string]*XmlMeta{}
	for _, file := range files {
		xml, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		WalkTree(xml, xmlMeta)
	}

	jsonBytes, err := json.MarshalIndent(xmlMeta, "", "    ")
	if err != nil {
		panic(err)
	}

	// WriteFile(name string, data []byte, perm FileMode) error
	err = os.WriteFile("/home/pewing/xml_support.xml", jsonBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// TODO: Temporary code, delete this
func WalkTree(bytes []byte, xmlMeta map[string]*XmlMeta) {
	doc := etree.NewDocument()

	err := doc.ReadFromBytes(bytes)
	if err != nil {
		panic(err)
	}

	walkTree(doc.Root(), xmlMeta)
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

func getElementPath(e *etree.Element) string {
	tags := []string{}

	for e != nil {
		tags = append(tags, e.Tag)
		e = e.Parent()
	}

	result := ""
	if len(tags) > 0 {
		result = tags[len(tags)-1]

		for i := len(tags) - 2; i >= 0; i-- {
			result = path.Join(result, tags[i])
		}
	}
	return result
}
