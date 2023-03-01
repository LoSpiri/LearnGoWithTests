package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("LoSpiri")
	want := "Hello LoSpiri"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
