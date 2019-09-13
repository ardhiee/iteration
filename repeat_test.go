package iteration

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t *testing.T, expected, repeated string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat(t *testing.T) {

	t.Run("repeated 'a' 5 times if zero inputed", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("repeated 'a' 7 times ", func(t *testing.T) {
		repeated := Repeat("a", 7)
		expected := "aaaaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}

func ExampleRepeat() {
	char := Repeat("b", 8)
	fmt.Println(char)
	// Output: bbbbbbbb
}
