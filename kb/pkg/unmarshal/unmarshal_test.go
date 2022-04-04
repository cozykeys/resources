package unmarshal

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseKeyboard(t *testing.T) {
	testCases := map[string]struct {
		dataFile string
	}{
		"happy path simple": {
			dataFile: "happy_path_simple.xml",
		},
		"bloomer_v2": {
			dataFile: "bloomer_v2.xml",
		},
		"bloomer_v3": {
			dataFile: "bloomer_v3.xml",
		},
		/*
		 */
		"bloomer_v4": {
			dataFile: "bloomer_v4.xml",
		},
		/*
			"speedo_v3": {
				dataFile: "speedo_v3.xml",
			},
			"triad_center": {
				dataFile: "triad_center.xml",
			},
			"triad_left": {
				dataFile: "triad_left.xml",
			},
			"triad_right": {
				dataFile: "triad_right.xml",
			},
		*/
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			xml, err := GetTestData(testCase.dataFile)
			require.Nil(t, err)

			kb, err := Unmarshal(xml)
			require.Nil(t, err)
			require.NotNil(t, kb)

			/*
				require.Equal(t, kb, &models.Keyboard{
					Name:      "bloomer",
					Version:   "1.0.0",
					Constants: nil,
					Layers:    nil,
				})
			*/
		})
	}
}

// TODO: Temporary, delete this
/*
func Test_WalkFiles(t *testing.T) {
	testDataDir := GetTestDataDirectory()

	files := []string{
		filepath.Join(testDataDir, "bloomer_v2.xml"),
		filepath.Join(testDataDir, "bloomer_v3.xml"),
		filepath.Join(testDataDir, "bloomer_v4.xml"),
		filepath.Join(testDataDir, "speedo_v3.xml"),
		filepath.Join(testDataDir, "triad_center.xml"),
		filepath.Join(testDataDir, "triad_left.xml"),
		filepath.Join(testDataDir, "triad_right.xml"),
	}

	WalkFiles(files)
}

func Test_WalkTree(t *testing.T) {

	//func (kb *Keyboard) WalkTree(bytes []byte) {

	xml, err := GetTestData("bloomer.xml")
	require.Nil(t, err)

	WalkTree(xml)
}
*/

func GetTestDataDirectory() string {
	_, testFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(testFile), "test_data")
}

func GetTestData(filename string) ([]byte, error) {
	testDataDirectory := GetTestDataDirectory()
	return os.ReadFile(filepath.Join(testDataDirectory, filename))
}
