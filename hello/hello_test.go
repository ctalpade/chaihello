package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world from Chaitanya Talpade"
	if got := SayHello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
