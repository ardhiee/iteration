package iteration

const repeatCount = 5

// Repeat func takes one character and return 5 character
func Repeat(character string, repeatCharacter int) string {

	var repeated string

	if repeatCharacter == 0 {
		repeatCharacter = repeatCount
	}

	for i := 0; i < repeatCharacter; i++ {
		repeated += character
	}
	return repeated
}
