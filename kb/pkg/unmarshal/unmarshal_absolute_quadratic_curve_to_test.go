package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalAbsoluteQuadraticCurveTo(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.AbsoluteQuadraticCurveTo
	}{
		"happy path": {
			xml: []byte(`<AbsoluteQuadraticCurveTo> <EndPoint X="1.0" Y="2.0" /> <ControlPoint X="-3.0" Y="4.0" /> </AbsoluteQuadraticCurveTo>`),
			expected: func() *models.AbsoluteQuadraticCurveTo {
				absoluteQuadraticCurveTo := &models.AbsoluteQuadraticCurveTo{
					KeyboardElementBase: models.KeyboardElementBase{
						Visible: true,
					},
				}
				absoluteQuadraticCurveTo.EndPoint = &models.Point{
					KeyboardElementBase: models.KeyboardElementBase{
						Parent:  absoluteQuadraticCurveTo,
						Visible: true,
					},
					X: 1.0,
					Y: 2.0,
				}
				absoluteQuadraticCurveTo.ControlPoint = &models.Point{
					KeyboardElementBase: models.KeyboardElementBase{
						Parent:  absoluteQuadraticCurveTo,
						Visible: true,
					},
					X: -3.0,
					Y: 4.0,
				}
				return absoluteQuadraticCurveTo
			}(),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			absoluteQuadraticCurveTo, err := unmarshalAbsoluteQuadraticCurveTo(doc.Root(), nil)
			require.Nil(t, err)
			require.Equal(t, testCase.expected, absoluteQuadraticCurveTo)
		})
	}
}
