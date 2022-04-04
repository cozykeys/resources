package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalCircle(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Circle
	}{
		"happy_path": {
			xml: []byte(`<Circle
							Size="5.0"
							XOffset="1.0"
							YOffset="2.0"
							Fill="#ffffff"
							Stroke="#000000"
							StrokeWidth="2px" />`),
			expected: &models.Circle{
				Size:        5.0,
				XOffset:     1.0,
				YOffset:     2.0,
				Fill:        "#ffffff",
				Stroke:      "#000000",
				StrokeWidth: "2px",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			circle, err := unmarshalCircle(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, circle)
		})
	}
}
