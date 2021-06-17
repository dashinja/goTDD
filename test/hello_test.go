package test

import (
	"testing"

	"github.com/dashinja/learn-go-with-tests/starter"
)

func TestHello(t *testing.T) {
	
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := starter.Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := starter.Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("if language parameter is empty, default to 'English'", func(t *testing.T) {
		got := starter.Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola' if the user passes in the language as 'Spanish'", func(t *testing.T) {

		got := starter.Hello("Diego", "Spanish")
		want := "Hola, Diego"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Bonjour' if language paramer is 'French'", func(t *testing.T) {
		got := starter.Hello("Pierre", "French")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say, 'Ni hao' if language paramer is 'Mandarin'", func(t *testing.T) {
		got := starter.Hello("Byron", "Mandarin")
		want := "Ni hao, Byron"
		assertCorrectMessage(t, got, want)
	})
}
