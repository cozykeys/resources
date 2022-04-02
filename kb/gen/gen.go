package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
	"unicode"
)

type unmarshalFileTemplateData struct {
	Name          string
	NameCamelCase string
	NameSnakeCase string
	Attributes    []string
	Children      []string
}

func main() {
	filename := "xml_support.json"
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}

	for k, v := range result {
		log.Printf("Key = %q, Value = %q", k, v)

		data := &unmarshalFileTemplateData{
			Name:          k,
			NameCamelCase: toCamelCase(k),
			NameSnakeCase: toSnakeCase(k),
			Attributes:    []string{},
		}

		attributes := v.(map[string]interface{})["Attributes"].([]interface{})
		for _, attr := range attributes {
			data.Attributes = append(data.Attributes, attr.(string))
		}

		children := v.(map[string]interface{})["Children"].([]interface{})
		for _, child := range children {
			data.Children = append(data.Attributes, child.(string))
		}

		{
			t, err := template.
				New(fmt.Sprintf("%s_impl", data.NameSnakeCase)).
				Parse(fileTemplateUnmarshalImpl)
			if err != nil {
				panic(err)
			}

			var b bytes.Buffer
			err = t.Execute(&b, data)
			if err != nil {
				panic(err)
			}

			filename := fmt.Sprintf("unmarshal_%s.go", data.NameSnakeCase)
			os.WriteFile(path.Join("temp", filename), b.Bytes(), 0644)
		}

		{
			t, err := template.
				New(fmt.Sprintf("%s_test", data.NameSnakeCase)).
				Parse(fileTemplateUnmarshalTest)
			if err != nil {
				panic(err)
			}

			var b bytes.Buffer
			err = t.Execute(&b, data)
			if err != nil {
				panic(err)
			}

			filename := fmt.Sprintf("unmarshal_%s_test.go", data.NameSnakeCase)
			os.WriteFile(path.Join("temp", filename), b.Bytes(), 0644)
		}
	}

	os.Exit(0)
}

func toCamelCase(s string) string {
	runes := []rune(s)
	if len(runes) > 0 && unicode.IsLetter(runes[0]) {
		lower := unicode.ToLower(runes[0])
		runes[0] = lower
	}
	return string(runes)
}

func toSnakeCase(s string) string {
	result := []rune{}
	for i, r := range []rune(s) {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, []rune{'_', unicode.ToLower(r)}...)
			} else {
				result = append(result, unicode.ToLower(r))
			}
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

var fileTemplateUnmarshalImpl = `package unmarshal

import (
	"kb/pkg/models"

	"github.com/beevik/etree"
)

func unmarshal{{.Name}}(e *etree.Element) (*models.{{.Name}}, error) {
	if e == nil {
		return nil, &nilElementError{}
	}

	if e.Tag != Element{{.Name}} {
		return nil, &invalidTagError{
			expected: Element{{.Name}},
			actual:   e.Tag,
		}
	}

	{{.NameCamelCase}} := &models.{{.Name}}{}

	err := unmarshal{{.Name}}Attributes({{.NameCamelCase}}, e.Attr)
	if err != nil {
		return nil, err
	}

	err = unmarshal{{.Name}}Children({{.NameCamelCase}}, e.Child)
	if err != nil {
		return nil, err
	}

	return {{.NameCamelCase}}, nil
}

func unmarshal{{.Name}}Attributes({{.NameCamelCase}} *models.{{.Name}}, attributes []etree.Attr) error {
	supportedAttributes := map[string]*struct {
		required bool
		found    bool
	}{ {{range .Attributes}}
		Attribute{{.}}: {required: false},{{end}}
	}

	for _, attr := range attributes {
		var err error
		switch attr.Key { {{range .Attributes}}
		case Attribute{{.}}:
			{{$.NameCamelCase}}.Name, err = unmarshalAttributeString(attr.Key, attr.Value){{end}}
		default:
			err = &unexpectedAttributeError{
				element:   Element{{.Name}},
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
				element:   Element{{.Name}},
				attribute: k,
			}
		}
	}

	return nil
}

func unmarshal{{.Name}}Children({{.NameCamelCase}} *models.{{.Name}}, children []etree.Token) error {
	for _, child := range children {
		element, ok := child.(*etree.Element)
		if !ok {
			// Ignore children that are not elements
			continue
		}

		var err error
		switch element.Tag {
		case ElementConstants:
			{{.NameCamelCase}}.Constants, err = unmarshalConstants(element)
		case ElementLayers:
			{{.NameCamelCase}}.Layers, err = unmarshalLayers(element)
		default:
			err = &invalidChildElementError{
				element: Element{{.Name}},
				child:   element.Tag,
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
`

var fileTemplateUnmarshalTest = `package unmarshal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_unmarshal{{.Name}}(t *testing.T) {
	require.Equal(t, 0, 0)
}
`
