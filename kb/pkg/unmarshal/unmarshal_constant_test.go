package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalConstant(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Constant
	}{
		"happy path": {
			xml: []byte(`<Constant Name="Foo" Value="Bar" />`),
			expected: &models.Constant{
				Name:  "Foo",
				Value: "Bar",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			constant, err := unmarshalConstant(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, constant)
		})
	}
}

func Test_unmarshalConstants(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected []models.Constant
	}{
		"happy path": {
			xml: []byte(`<Constants><Constant Name="Foo" Value="Bar" /></Constants>`),
			expected: []models.Constant{
				{
					Name:  "Foo",
					Value: "Bar",
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			constants, err := unmarshalConstants(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, constants)
		})
	}
}
