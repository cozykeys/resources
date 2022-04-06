package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalSpacer(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Spacer
	}{
		"happy_path": {
			xml: []byte(`<Spacer Height="1.0" Width="2.0" />`),
			expected: &models.Spacer{
				KeyboardElementBase: models.KeyboardElementBase{
					Height:  1.0,
					Width:   2.0,
					Visible: true,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			spacer, err := unmarshalSpacer(doc.Root(), nil)
			require.Nil(t, err)
			require.Equal(t, testCase.expected, spacer)
		})
	}
}
