package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeated("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeated("a")
	}
}
