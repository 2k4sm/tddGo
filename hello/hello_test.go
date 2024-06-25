package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello normally", func(t *testing.T) {
		got := Hello("shrinibas")
		want := "Hello, shrinibas
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
