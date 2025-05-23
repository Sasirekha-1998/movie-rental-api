package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello"
	got := "Sasi"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
