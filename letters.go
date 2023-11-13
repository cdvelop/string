package strings

var VALID_LETTERS = map[byte]byte{
	'A': 'a', 'B': 'b', 'C': 'c', 'D': 'd', 'E': 'e', 'F': 'f', 'G': 'g', 'H': 'h', 'I': 'i',
	'J': 'j', 'K': 'k', 'L': 'l', 'M': 'm', 'N': 'n', 'O': 'o', 'P': 'p', 'Q': 'q', 'R': 'r',
	'S': 's', 'T': 't', 'U': 'u', 'V': 'v', 'W': 'w', 'X': 'x', 'Y': 'y', 'Z': 'z',
}

func Letters() map[rune]string {
	return map[rune]string{
		'A': "a", 'B': "b", 'C': "c", 'D': "d", 'E': "e", 'F': "f", 'G': "g", 'H': "h", 'I': "i",
		'J': "j", 'K': "k", 'L': "l", 'M': "m", 'N': "n", 'O': "o", 'P': "p", 'Q': "q", 'R': "r",
		'S': "s", 'T': "t", 'U': "u", 'V': "v", 'W': "w", 'X': "x", 'Y': "y", 'Z': "z",
		'Ñ': "ñ",
	}
}
