package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalAbsoluteLineTo(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.AbsoluteLineTo
	}{
		"happy path": {
			xml: []byte(`<AbsoluteLineTo><EndPoint X="1.0" Y="2.0" /></AbsoluteLineTo>`),
			expected: func() *models.AbsoluteLineTo {
				absoluteLineTo := &models.AbsoluteLineTo{}
				absoluteLineTo.EndPoint = &models.Point{
					KeyboardElementBase: models.KeyboardElementBase{
						Parent: absoluteLineTo,
					},
					X: 1.0,
					Y: 2.0,
				}
				return absoluteLineTo
			}(),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			absoluteLineTo, err := unmarshalAbsoluteLineTo(doc.Root(), nil)
			require.Nil(t, err)
			require.Equal(t, testCase.expected, absoluteLineTo)
		})
	}
}
