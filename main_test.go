package main

import "testing"

func TestSayHello(t *testing.T) {
	expected := "Hello, World!"
	if sayHello() != expected {
		t.Fatalf("expected %s but got %s", expected, sayHello())
	}
}
