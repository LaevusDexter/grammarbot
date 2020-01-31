package grammarbot

// CorrectMatchesBytes - returns corrected text with first replacements in matches
func CorrectMatchesBytes(text []byte, matches []*Match) []byte {
	var (
		i int
		offset int
		extraSize int
	)

	for i = 0; i < len(matches); i++ {
		extraSize += len(matches[i].Replacements[0].Value) - matches[i].Length
	}

	buf := make([]byte, len(text) + extraSize)

	offset = 0
	currentPos := 0
	nextPos := 0
	for i = 0; i < len(matches); i++ {
		switch i + 1 {
		case 1:
			currentPos = 0
			nextPos = matches[i].Offset
		default:
			currentPos = matches[i - 1].Offset +  matches[i - 1].Length
			nextPos = matches[i].Offset
		}

		offset += copy(buf[offset:], text[currentPos:nextPos])
		offset += copy(buf[offset:], matches[i].Replacements[0].Value)
	}

	currentPos = matches[i - 1].Offset  +  matches[i - 1].Length
	nextPos = len(text)
	copy(buf[offset:], text[currentPos:nextPos])

	return buf
}

// CorrectMatches - returns corrected text with first replacements in matches
func CorrectMatches(text string, matches []*Match) string {
	/*
		s2b cast creates undefined behavior (runtime might ask for cap), so here's copy/paste from FixMatchesBytes
	*/

	var (
		i int
		offset int
		extraSize int
	)

	for i = 0; i < len(matches); i++ {
		extraSize += len(matches[i].Replacements[0].Value) - matches[i].Length
	}

	buf := make([]byte, len(text) + extraSize)

	offset = 0
	currentPos := 0
	nextPos := 0
	for i = 0; i < len(matches); i++ {
		switch i + 1 {
		case 1:
			currentPos = 0
			nextPos = matches[i].Offset
		default:
			currentPos = matches[i - 1].Offset +  matches[i - 1].Length
			nextPos = matches[i].Offset
		}

		offset += copy(buf[offset:], text[currentPos:nextPos])
		offset += copy(buf[offset:], matches[i].Replacements[0].Value)
	}

	currentPos = matches[i - 1].Offset  +  matches[i - 1].Length
	nextPos = len(text)
	copy(buf[offset:], text[currentPos:nextPos])

	return string(buf)
}