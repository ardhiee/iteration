package iteration

// Repeated only return character
func Repeated(character string, repeatCount int) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
