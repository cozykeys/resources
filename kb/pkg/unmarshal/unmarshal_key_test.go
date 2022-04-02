package unmarshal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_unmarshalKey(t *testing.T) {
	testCases := map[string]struct {
		xml []byte
	}{
		"happy path": {
			xml: []byte("TODO"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, []byte("TODO"), testCase.xml)
		})
	}
}
