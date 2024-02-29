package strings

func (Strings) IsNotLetter(r rune) bool {
	return (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') &&
		(r != 'Ñ' && r != 'ñ') &&
		(r != 'Á' && r != 'É' && r != 'Í' && r != 'Ó' && r != 'Ú' &&
			r != 'á' && r != 'é' && r != 'í' && r != 'ó' && r != 'ú')
}
