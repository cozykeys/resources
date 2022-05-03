package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalPoint(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Point
	}{
		"happy_path_end_point": {
			xml: []byte(`<EndPoint X="1.0" Y="2.0" />`),
			expected: &models.Point{
				KeyboardElementBase: models.KeyboardElementBase{
					Visible: true,
				},
				X: 1.0,
				Y: 2.0,
			},
		},
		"happy_path_control_point": {
			xml: []byte(`<ControlPoint X="3.0" Y="4.0" />`),
			expected: &models.Point{
				KeyboardElementBase: models.KeyboardElementBase{
					Visible: true,
				},
				X: 3.0,
				Y: 4.0,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			point, err := unmarshalPoint(doc.Root(), nil)
			require.Nil(t, err)
			require.Equal(t, testCase.expected, point)
		})
	}
}
