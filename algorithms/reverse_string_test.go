package algorithms

import "testing"

func TestReverseString(t *testing.T) {
    got := ReverseString("hello")
    want := "olleh"
    if got != want {
        t.Errorf("ReverseString(\"hello\") = %q, want %q", got, want)
    }
}