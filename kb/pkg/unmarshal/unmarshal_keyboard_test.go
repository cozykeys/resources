package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalKeyboard(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Keyboard
	}{
		"happy path": {
			xml: []byte(`<Keyboard Name="Bloomer" Version="1.0.0"><Constants /><Layers /></Keyboard>`),
			expected: &models.Keyboard{
				Name:      "Bloomer",
				Version:   "1.0.0",
				Constants: []models.Constant{},
				Layers:    []models.Layer{},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			keyboard, err := unmarshalKeyboard(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, keyboard)
		})
	}
}
