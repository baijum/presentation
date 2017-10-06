package hello

import "testing"

func TestSayHello(t *testing.T) {
	out := SayHello()
	expected := "Hello"
	if len(out) != 6 {
		t.FailNow()
	}
	if expected != out {
		t.Error("output not matching")
	}
}
func TestSayHello2(t *testing.T) {
	out := SayHello()
	expected := "Hello"
	if len(out) != 6 {
		t.Error("length is not matching")
	}
	if expected != out {
		t.Error("output not matching")
	}
}
