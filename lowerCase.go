package strings

// todo a mayúscula alfabeto con ñ
func (s Strings) ToUpperCase(text string) string {
	return s.upperLowerCase(&text, true, false)
}

// todo a minúscula alfabeto con ñ
func (s Strings) ToLowerCase(text string) string {
	return s.upperLowerCase(&text, false, false)
}

func (s Strings) UpperCaseFirstLetter(text string) string {
	return s.upperLowerCase(&text, true, true)
}

func (s Strings) LowerCaseFirstLetter(text string) string {
	return s.upperLowerCase(&text, false, true)
}

func (s Strings) upperLowerCase(text *string, upper, onlyFirstLetter bool) (out string) {

	if len(*text) == 0 {
		return *text
	}

	// las runas puede tener mas de 2 caracteres por eso no es conveniente usar text[0]
	for i, r := range *text {

		char := string(r)

		if s.IsNotLetter(r) {
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
