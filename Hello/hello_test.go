package Hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Tanawat", "")
		want := "Hello, Tanawat"
		assertCorrectMessage(t, got, want)
	})

	t.Run(`Say "Hello, World" when an empty string is supplied`, func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run(`Say hello thai language`, func(t *testing.T) {
		got := Hello("คนึงนิจ", "thai")
		want := "สวัสดี, คนึงนิจ"
		assertCorrectMessage(t, got, want)
	})

	t.Run(`Say hello french language`, func(t *testing.T) {
		got := Hello("Mike", "french")
		want := "Salut, Mike"
		assertCorrectMessage(t, got, want)
	})
}
