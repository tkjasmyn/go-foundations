package algorithms

import "testing"

func TestFizzBuzz(t *testing.T) {
    got := FizzBuzz()
    if len(got) != 100 {
        t.Errorf("FizzBuzz() returned %d items, want 100", len(got))
    }
    if got[2] != "Fizz" {
        t.Errorf("got[2] = %q, want \"Fizz\"", got[2])
    }
    if got[4] != "Buzz" {
        t.Errorf("got[4] = %q, want \"Buzz\"", got[4])
    }
    if got[14] != "FizzBuzz" {
        t.Errorf("got[14] = %q, want \"FizzBuzz\"", got[14])
    }
}