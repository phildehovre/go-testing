package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Will say hello <name>", func(t *testing.T) {
		got := SayHello("Phil", "")
		want := "Hello, Phil"

		assessSayHello(t, got, want)
	})

	t.Run("say hello world when no string is provided", func(t *testing.T) {

		got := SayHello("", "")
		want := "Hello, world"

		assessSayHello(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := SayHello("Phil", "Spanish")
		want := "Hola, Phil"
		assessSayHello(t, got, want)
	})
}

func assessSayHello(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
