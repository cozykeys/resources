package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalKey(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Key
	}{
		"happy_path": {
			xml: []byte(`<Key
							Name="k13"
							Row="1"
							Column="2"
							XOffset="-1.0"
							YOffset="3.0"
							Width="18.1"
							Height="18.1"
							Margin="0.475"
							Fill="#ffffff"
							Stroke="#000000">
							<Legend
								HorizontalAlignment="Center"
								VerticalAlignment="Center"
								Text="F1"
								FontSize="4"
								Color="#000000" />
							</Key>`),
			expected: func() *models.Key {
				key := &models.Key{
					KeyboardElementBase: models.KeyboardElementBase{
						Name:    "k13",
						XOffset: -1.0,
						YOffset: 3.0,
						Width:   18.1,
						Height:  18.1,
						Margin:  0.475,
						Visible: true,
					},
					Row:    1,
					Column: 2,
					Fill:   "#ffffff",
					Stroke: "#000000",
				}

				key.Legends = []models.Legend{
					{
						KeyboardElementBase: models.KeyboardElementBase{
							Parent:  key,
							Visible: true,
						},
						HorizontalAlignment: models.LegendHorizontalAlignmentCenter,
						VerticalAlignment:   models.LegendVerticalAlignmentCenter,
						Text:                "F1",
						FontSize:            4,
						Color:               "#000000",
					},
				}

				return key
			}(),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			key, err := unmarshalKey(doc.Root(), nil)
			require.Nil(t, err)
			require.Equal(t, testCase.expected, key)
		})
	}
}
