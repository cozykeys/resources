package models

func mergeConstants(parent, child []Constant) []Constant {
	kvps := map[string]string{}

	for _, constant := range parent {
		kvps[constant.Name] = constant.Value
	}

	for _, constant := range child {
		kvps[constant.Name] = constant.Value
	}

	result := make([]Constant, len(kvps))
	i := 0
	for k, v := range kvps {
		result[i] = Constant{Name: k, Value: v}
		i++
	}
	return result
}
