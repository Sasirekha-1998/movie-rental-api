package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello"
	got := "Hello"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
