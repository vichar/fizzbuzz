package fizzbuzz_test

import "testing"
import "github.com/vichar/fizzbuzz"

func TestFizzBuzz(t *testing.T) {
	t.Run("1 should say 1", func(t *testing.T) {
		get := fizzbuzz.Say(1)
		want := "1"
		if want != get {
			t.Errorf("It should say %s but get %s", want, get)
		}
	})
}
