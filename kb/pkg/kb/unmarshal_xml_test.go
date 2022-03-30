package kb

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseKeyboard(t *testing.T) {
	testCases := map[string]struct {
		dataFile string
	}{
		/*
			"happy path simple": {
				dataFile: "happy_path_simple.xml",
			},
		*/
		"bloomer": {
			dataFile: "bloomer.xml",
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			xml, err := GetTestData(testCase.dataFile)
			require.Nil(t, err)

			kb, err := ParseKeyboard(xml)
			require.Nil(t, err)
			require.NotNil(t, kb)
		})
	}
}

func GetTestData(filename string) ([]byte, error) {
	_, testFile, _, _ := runtime.Caller(0)

	testDataDirectory := filepath.Join(filepath.Dir(testFile), "test_data")
	testDataFile := filepath.Join(testDataDirectory, filename)

	return ioutil.ReadFile(testDataFile)
}
