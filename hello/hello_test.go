package hello

import ("fmt"
		"testing"
		)

func TestHello(t *testing.T) {
	want := "Hello, world from Chaitanya Talpade"
	if got := SayHello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestOneMore(t *testing.T) {
	want := "Hello, world from Chaitanya Talpade"
	if got := OneMore("Chai"); got == want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
	
	fmt.Println("OneMore func called")
}
