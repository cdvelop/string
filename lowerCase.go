package strings

// solo a minúscula texto del alfabeto con ñ
func ToLowerCase(in string) string {

	var out string

	for _, c := range in {

		if l, exist := LettersUpperLowerCase(true)[c]; exist {
			out += string(l)
		} else {
			out += string(c)
		}
	}

	return out
}

func LowerCaseFirstLetter(name string) string {
	if newChar, ok := LettersUpperLowerCase()[rune(name[0])]; ok {
		return string(newChar) + name[1:]
	}
	return name
}
