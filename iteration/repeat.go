package iteration

// Repeat returns a string that repeats the character by specified times.
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
