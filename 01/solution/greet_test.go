package main

import "testing"

func TestGreet_ByDefault_GreetsInput(t *testing.T) {
	expected := "Hello, Banana"
	actual := Greet("Banana")

	if expected != actual {
		t.Errorf("unexpected greeting. expected %v actual %v", expected, actual)
	}
}
