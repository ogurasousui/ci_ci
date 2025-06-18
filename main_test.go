package main

import "testing"

func TestSayHello(t *testing.T) {
	expected := "Hello, World!"
	if sayHello() != expected {
		t.Fatalf("expected %s but got %s", expected, sayHello())
	}
}

func TestAdd(t *testing.T) {
	sum := add(2, 3)
	if sum != 5 {
		t.Fatalf("expected 5 but got %d", sum)
	}
}
