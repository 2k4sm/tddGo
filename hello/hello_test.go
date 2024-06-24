package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("shrinibas")
	want := "Hello there shrinibas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
