package main

import "testing"

func TestMain(t *testing.T) {

	t.Run("say hello to someone", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assetCorrectMessage(t, got, want)
	})

	t.Run("say hello to the world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assetCorrectMessage(t, got, want)
	})
}

func assetCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
