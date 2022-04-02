package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalPath(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Path
	}{
		"happy path": {
			xml: []byte(`<Path Fill="#ffffff" FillOpacity="0.5" Stroke="#000000" StrokeWidth="1.0" Visible="True"><Components /></Path>`),
			expected: &models.Path{
				Fill:        "#ffffff",
				FillOpacity: "0.5",
				Stroke:      "#000000",
				StrokeWidth: "1.0",
				Visible:     true,
				Components:  []models.PathComponent{},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			path, err := unmarshalPath(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, path)
		})
	}
}
