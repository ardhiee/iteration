package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("will repeat 'a' five times", func(t *testing.T) {
		got := Repeated("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("will repeat 'a' 7 times", func(t *testing.T) {
		got := Repeated("a", 7)
		want := "aaaaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeated("a", 5)
	}
}
