package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	want := 4
	assertCorrectSum(t, sum, want)
}

func assertCorrectSum(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
