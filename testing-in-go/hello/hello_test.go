package hello

import "testing"

func TestSayHello(t *testing.T) {
	out := SayHello()
	expected := "Hello"
	if expected != out {
		t.Error("No Hello output")
	}
}
