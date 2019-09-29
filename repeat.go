package iteration

const repeatCount = 5

// Repeated only return character
func Repeated(character string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
