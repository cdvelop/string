package strings

var letters = map[rune]rune{
	'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i',
	'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r',
	'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z',
}

// ej: A:a, B:b.. with_ñ = Ñ:ñ
func LettersUpperLowerCase(with_ñ ...bool) map[rune]rune {

	for _, ñ := range with_ñ {
		if ñ {
			letters['Ñ'] = 'ñ'
			return letters
		}
	}

	// Si no se incluye la letra "Ñ", se elimina del mapa
	delete(letters, 'Ñ')

	return letters
}
