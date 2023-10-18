package string

import "github.com/cdvelop/model"

// solo a minúscula texto del alfabeto
func ToLowercaseAlphabet(in string) (string, error) {

	var out string

	for _, c := range in {

		if l, exist := Letters()[c]; exist {
			out += l
		} else if c == ' ' {
			out += " "
		} else {
			return "", model.Error("carácter:", string(c), "no soportado para convertir a minúscula")
		}
	}

	return out, nil
}
