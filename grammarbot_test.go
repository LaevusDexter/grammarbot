package grammarbot

import (
	"testing"
)

func TestCheckAndCorrect(t *testing.T) {
	text := `
		The man jump into a black sedan and he drived away before being noticed. 
		Also he can't remember how to get to thei wroom.`

	c, err := Check(text)
	if err != nil {
		t.Fatal(err)
	}

	text = CorrectMatches(text, c.Matches)

	const correctedText = `
		The man jumped into a black sedan, and he drove away before being noticed. 
		Also, he can't remember how to get to the room.`

	if text != correctedText {
		t.Fatal("\n", "Expected:", correctedText, "\n", "Got:", text)
	}

}