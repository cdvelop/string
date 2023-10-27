package strings

// ej: "texto1,texto2" "," = []string{texto1,texto2}
func Split(data, separator string) (result []string) {

	if len(data) >= 3 {

		start := 0

		for i := 0; i < len(data); i++ {
			if data[i:i+len(separator)] == separator {
				result = append(result, data[start:i])
				start = i + len(separator)
				i += len(separator) - 1
			}
		}

		result = append(result, data[start:])
	} else {
		return []string{data}
	}

	return
}
