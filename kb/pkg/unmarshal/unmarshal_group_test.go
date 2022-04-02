package unmarshal

import (
	"kb/pkg/models"
	"testing"

	"github.com/beevik/etree"
	"github.com/stretchr/testify/require"
)

func Test_unmarshalGroup(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected *models.Group
	}{
		"happy path": {
			xml: []byte(`<Group Name="ClusterLeft" Rotation="10" XOffset="-93.331" YOffset="-4.815" Visible="true"><Children /></Group>`),
			expected: &models.Group{
				Name:     "ClusterLeft",
				Rotation: 10.0,
				XOffset:  -93.331,
				YOffset:  -4.815,
				Visible:  true,
				Children: []models.GroupChild{},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			group, err := unmarshalGroup(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, group)
		})
	}
}

func Test_unmarshalGroups(t *testing.T) {
	testCases := map[string]struct {
		xml      []byte
		expected []models.Group
	}{
		"happy path": {
			xml: []byte(`<Groups><Group Name="ClusterLeft" Rotation="10" XOffset="-93.331" YOffset="-4.815" Visible="true"><Children /></Group></Groups>`),
			expected: []models.Group{
				{
					Name:     "ClusterLeft",
					Rotation: 10.0,
					XOffset:  -93.331,
					YOffset:  -4.815,
					Visible:  true,
					Children: []models.GroupChild{},
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			doc := etree.NewDocument()

			err := doc.ReadFromBytes(testCase.xml)
			require.Nil(t, err)

			groups, err := unmarshalGroups(doc.Root())
			require.Nil(t, err)
			require.Equal(t, testCase.expected, groups)
		})
	}
}
