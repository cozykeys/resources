package kb

import (
	"encoding/json"
	"fmt"
	"kb/pkg/geo"
	"os"
	"strconv"
)

func GetInputVertices(filepath string) ([]*geo.Point3D, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read input file: %v", err)
	}

	var points []*geo.Point3D
	if err := json.Unmarshal(bytes, &points); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input vertices: %v", err)
	}

	return points, nil
}

func MustMarshalJSON(i interface{}, indent bool) []byte {
	var err error
	var bytes []byte
	if indent {
		bytes, err = json.MarshalIndent(i, "", "  ")
	} else {
		bytes, err = json.Marshal(i)
	}
	if err != nil {
		panic(err)
	}
	return bytes
}

func MustParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse float64: %v", err))
	}
	return f
}
