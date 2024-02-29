package strings

var upperLetters = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'Ñ', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
var lowerLetters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'ñ', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

// todo a mayúscula alfabeto con ñ
func ToUpperCase(text string) string {
	return upperLowerCase(&text, true, false)
}

// todo a minúscula alfabeto con ñ
func ToLowerCase(text string) string {
	return upperLowerCase(&text, false, false)
}

func UpperCaseFirstLetter(text string) string {
	return upperLowerCase(&text, true, true)
}

func LowerCaseFirstLetter(text string) string {
	return upperLowerCase(&text, false, true)
}

func upperLowerCase(text *string, upper, onlyFirstLetter bool) (out string) {

	if len(*text) == 0 {
		return *text
	}

	// las runas puede tener mas de 2 caracteres por eso no es conveniente usar text[0]
	for i, r := range *text {

		char := string(r)

		if IsNotLetter(r) {
			out += char // Agregar directamente a la salida
			continue
		}

		for index, upLetter := range upperLetters {

			// fmt.Printf("letter: %c lower: %v \n", r, string(lowerLetters[index]))
			if upper {
				if r == lowerLetters[index] {
					char = string(upLetter)
					break
				}
			} else { // down

				if r == upLetter {
					char = string(lowerLetters[index])
					break
				}

			}

		}

		if onlyFirstLetter && i == 0 { // retornamos inmediatamente
			return char + (*text)[len(string(r)):] //primera letra + tamaño de la runa mas el resto del texto
		}

		out += char // salida

		// fmt.Printf("actual: [%v] letra: [%v]\n", out, char)
	}

	return
}

func IsNotLetter(r rune) bool {
	return (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r != 'Ñ' && r != 'ñ')
}
