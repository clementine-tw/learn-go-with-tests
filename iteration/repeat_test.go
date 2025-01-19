package iteration

import "testing"

func TestRepeat(t *testing.T) {
	character := "a"
	count := 5
	repeated := Repeat(character, count)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected %q when repeat %q by count %d, but got %q",
			expected, character, count, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
