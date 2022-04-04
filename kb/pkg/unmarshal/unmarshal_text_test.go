package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalText(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Text
	}{
		"happy_path": {
			xml: []byte(`<Text 
                            Content="Foo Bar"
                            TextAnchor="middle"
                            Font="3px sans-serif"
                            Fill="#ffff00"
                            XOffset="1.0"
                            YOffset="2.0" />`),
			expected: &models.Text{
				Content:    "Foo Bar",
				TextAnchor: "middle",
				Font:       "3px sans-serif",
				Fill:       "#ffff00",
				XOffset:    1.0,
				YOffset:    2.0,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			text, err := unmarshalText(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, text)
		})
	}
}
