package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalAbsoluteMoveTo(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.AbsoluteMoveTo
	}{
		"happy path": {
			xml: []byte(`<AbsoluteMoveTo><EndPoint X="1.0" Y="2.0" /></AbsoluteMoveTo>`),
			expected: func() *models.AbsoluteMoveTo {
				absoluteMoveTo := &models.AbsoluteMoveTo{}
				absoluteMoveTo.EndPoint = &models.Point{
					KeyboardElementBase: models.KeyboardElementBase{
						Parent: absoluteMoveTo,
					},
					X: 1.0,
					Y: 2.0,
				}
				return absoluteMoveTo
			}(),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			absoluteMoveTo, err := unmarshalAbsoluteMoveTo(doc.Root(), nil)
			require.Nil(t, err)
			require.Equal(t, testCase.expected, absoluteMoveTo)
		})
	}
}
