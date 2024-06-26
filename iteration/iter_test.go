package iteration

import "testing"

func TestIter(t *testing.T) {
	t.Run("test duplicating a string", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "aaa"

		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})

	t.Run("test duplicating a string `k` times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("expected %q, got %q", want, got)
		}
	})
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
