package runeByIndex

func RuneByIndex(s *string, i *int) (ch rune, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	} ()

	return []rune(*s)[*i], nil
}
