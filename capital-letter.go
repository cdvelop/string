package strings

func (s Strings) FirstLetterIsCapital(text string) bool {

	for _, r := range text {

		if s.IsNotLetter(r) {
			continue
		}

		for _, upLetter := range upperLetters {

			if r == upLetter {
				return true
			}

		}

		return false

	}

	return false
}
