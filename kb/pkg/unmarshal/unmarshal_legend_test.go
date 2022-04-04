package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalLegend(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Legend
	}{
		"happy_path": {
			xml: []byte(`<Legend 
                            HorizontalAlignment="Center"
                            VerticalAlignment="Center"
                            Text="+"
                            FontSize="6"
                            Color="#000000"
                            YOffset="-3" />`),
			expected: &models.Legend{
				HorizontalAlignment: models.LegendHorizontalAlignmentCenter,
				VerticalAlignment:   models.LegendVerticalAlignmentCenter,
				Text:                "+",
				FontSize:            6.0,
				Color:               "#000000",
				YOffset:             -3.0,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			legend, err := unmarshalLegend(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, legend)
		})
	}
}
