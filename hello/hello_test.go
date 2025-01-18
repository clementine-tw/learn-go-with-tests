package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string applied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// It is a good idea to use testing.TB, which can access test or benchmark,
	// in a helper function.

	// The t.Helper() is needed to tell tester that this function is only a helper,
	// and track the actual line which error occured.
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
