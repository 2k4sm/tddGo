package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello normally", func(t *testing.T) {
		got := Hello("shrinibas", "english")
		want := "Hello, shrinibas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to empty string", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test spanish hello", func(t *testing.T) {
		got := Hello("prem", "spanish")
		want := "Hola, prem"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test french hello", func(t *testing.T) {
		got := Hello("bhav", "french")
		want := "Bonjour, bhav"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test random hello", func(t *testing.T) {
		got := Hello("bhav", "")
		want := "Hello, bhav"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
