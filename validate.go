package strings

func (Strings) IsNotLetter(r rune) bool {
	return (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r != 'Ñ' && r != 'ñ')
}
