package strings

// solo a minúscula texto del alfabeto
func ToLowerCaseAlphabet(in string) string {

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
