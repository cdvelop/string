package strings

// solo a minúscula texto del alfabeto
func ToLowerCase(in string) string {

	var out string

	for _, c := range in {

		if l, exist := Letters()[c]; exist {
			out += l
		} else {
			out += string(c)
		}
	}

	return out
}

func LowerCaseFirstLetter(name string) string {
	if newChar, ok := VALID_LETTERS[name[0]]; ok {
		return string(newChar) + name[1:]
	}
	return name
}
