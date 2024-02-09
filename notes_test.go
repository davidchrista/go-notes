package main

import "testing"

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
	result = Fooer(4)
	if result != "4" {
		t.Errorf("Result was incorrect, got %s, want: %s.", result, "4")
	}
}
