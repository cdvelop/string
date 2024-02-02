package strings

func Replace(text, old, newStr string) (result string) {

	for i := 0; i < len(text); i++ {
		// Buscar la ocurrencia de la palabra antigua en el texto
		if i+len(old) <= len(text) && text[i:i+len(old)] == old {
			// Agregar la nueva palabra al resultado
			result += newStr
			// Saltar la longitud de la palabra antigua en el texto original
			i += len(old) - 1
		} else {
			// Agregar el carácter actual al resultado
			result += string(text[i])
		}
	}

	return result
}

func TrimSuffix(text, suffix string) string {
	if len(text) < len(suffix) || text[len(text)-len(suffix):] != suffix {
		return text
	}
	return text[:len(text)-len(suffix)]
}

// Eliminar espacios al principio y al final
func Trim(text string) string {
	// Eliminar espacios al principio
	start := 0
	for start < len(text) && text[start] == ' ' {
		start++
	}

	// Eliminar espacios al final y al final de cada línea
	end := len(text) - 1
	for end >= 0 && (text[end] == ' ' || text[end] == '\n' || text[end] == '\t') {
		end--
	}

	// Caso especial: cadena vacía
	if start > end {
		return ""
	}

	// Devolver la subcadena sin espacios
	return text[start : end+1]
}
