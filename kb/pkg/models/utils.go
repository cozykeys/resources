package models

func mergeConstants(parent, child map[string]string) map[string]string {
	kvps := map[string]string{}

	for name, value := range parent {
		kvps[name] = value
	}

	for name, value := range child {
		kvps[name] = value
	}

	return kvps
}
