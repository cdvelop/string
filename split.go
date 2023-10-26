package strings

func Split(data, separator string) ([]string, error) {
	var result []string
	start := 0

	for i := 0; i < len(data); i++ {
		if data[i:i+len(separator)] == separator {
			result = append(result, data[start:i])
			start = i + len(separator)
			i += len(separator) - 1
		}
	}

	result = append(result, data[start:])

	return result, nil
}
