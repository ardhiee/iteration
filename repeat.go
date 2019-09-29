package iteration

// Repeated only return character
func Repeated(character string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated = repeated + "a"
	}
	return repeated
}
