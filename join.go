package strings

// separate ej: " ", "," "."
func Join(values []string, separate string) (out string) {

	var space string
	for _, v := range values {

		out += space + v
		space = separate

	}
	return
}
