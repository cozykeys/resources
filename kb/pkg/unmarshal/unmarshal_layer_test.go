package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalLayer(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Layer
	}{
		"happy path": {
			xml: []byte(`<Layer Name="Foo" ZIndex="1" XOffset="1.0" YOffset="2.0"><Groups /></Layer>`),
			expected: &models.Layer{
				KeyboardElementBase: models.KeyboardElementBase{
					Name:    "Foo",
					XOffset: 1.0,
					YOffset: 2.0,
					Visible: true,
				},
				ZIndex: 1,
				Groups: []models.Group{},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			layer, err := unmarshalLayer(doc.Root(), nil)
			require.Nil(t, err)
			require.Equal(t, testCase.expected, layer)
		})
	}
}

func Test_unmarshalLayers(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected []models.Layer
	}{
		"happy path": {
			xml: []byte(`<Layers><Layer Name="Foo" ZIndex="1" XOffset="1.0" YOffset="2.0"><Groups /></Layer></Layers>`),
			expected: []models.Layer{
				{
					KeyboardElementBase: models.KeyboardElementBase{
						Name:    "Foo",
						XOffset: 1.0,
						YOffset: 2.0,
						Visible: true,
					},
					ZIndex: 1,
					Groups: []models.Group{},
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			layers, err := unmarshalLayers(doc.Root(), nil)
			require.Nil(t, err)
			require.Equal(t, testCase.expected, layers)
		})
	}
}
